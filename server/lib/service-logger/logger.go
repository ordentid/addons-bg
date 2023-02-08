package servicelogger

import (
	"os"
	"strconv"
	"strings"

	"github.com/evalphobia/logrus_fluent"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type customLogger struct {
	ecslogrus   *ecslogrus.Formatter
	serviceName string
	hostName    string
	formatter   logrus.Formatter
}

func newCustomFormatter(serviceName string, hostname string) *customLogger {

	return &customLogger{
		ecslogrus: &ecslogrus.Formatter{
			DataKey:     "data_details",
			PrettyPrint: true,
		},
		serviceName: serviceName,
		hostName:    hostname,
		formatter:   logrus.StandardLogger().Formatter,
	}
}

func (l *customLogger) Format(entry *logrus.Entry) ([]byte, error) {
	// set ecs format
	entry.Data["service_name"] = l.serviceName
	entry.Data["host_name"] = l.hostName

	return l.ecslogrus.Format(entry)
}

func New(serviceName string) *logrus.Logger {
	name := strings.ToLower(serviceName)
	name = strings.ReplaceAll(name, " ", "_")

	output := getEnv("LOG_OUTPUT", "")
	host := getEnv("FLUENTBIT_HOST", "localhost")
	sport := getEnv("FLUENTBIT_PORT", "24223")
	port, _ := strconv.Atoi(sport)

	hostname, _ := os.Hostname()

	log := logrus.New()

	log.SetFormatter(newCustomFormatter(name, hostname))

	// Add the GlobalKeyHook to the logrus.Logger
	log.Hooks.Add(&GlobalKeyHook{
		keys: logrus.Fields{
			"service_name": name,
			"host_name":    hostname,
		},
		esl: &ecslogrus.Formatter{
			DataKey: "data_details",
		},
	})

	if strings.ToLower(output) == "elastic" {

		hook, err := logrus_fluent.NewWithConfig(logrus_fluent.Config{
			Port: port,
			Host: host,
		})
		if err != nil {
			panic(err)
		}

		logrus.Info("no error connection for fluentbit")

		hook.SetLevels(logrus.AllLevels)
		hook.SetTag(name)

		hook.SetMessageField("message")

		// ignore field
		// hook.AddIgnore("context")

		// filter func
		hook.AddFilter("error", logrus_fluent.FilterError)

		log.AddHook(hook)

	}

	log.ReportCaller = true

	return log
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type errorObject struct {
	Message string `json:"message,omitempty"`
}

type GlobalKeyHook struct {
	keys logrus.Fields
	esl  *ecslogrus.Formatter
}

func (h *GlobalKeyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GlobalKeyHook) Fire(entry *logrus.Entry) error {
	// now := time.Now().Format("2006-01-02T15:04:05.000Z0700")

	for k, v := range h.keys {
		entry.Data[k] = v
	}

	datahint := len(entry.Data)
	if h.esl.DataKey != "" {
		datahint = 2
	}
	data := make(logrus.Fields, datahint)
	if len(entry.Data) > 0 {
		extraData := data
		if h.esl.DataKey != "" {
			extraData = make(logrus.Fields, len(entry.Data))
		}
		for k, v := range entry.Data {
			switch k {
			case logrus.ErrorKey:
				err, ok := v.(error)
				if ok {
					data["error"] = errorObject{
						Message: err.Error(),
					}
					break
				}
				fallthrough // error has unexpected type
			default:
				if k != "service_name" && k != "data_tag" {
					delete(entry.Data, k)
				}
				extraData[k] = v
			}
		}
		if h.esl.DataKey != "" && len(extraData) > 0 {
			data[h.esl.DataKey] = extraData
		}
	}
	if entry.HasCaller() {
		// Logrus has a single configurable field (logrus.FieldKeyFile)
		// for storing a combined filename and line number, but we want
		// to split them apart into two fields. Remove the event's Caller
		// field, and encode the ECS fields explicitly.
		var funcVal, fileVal string
		var lineVal int
		if h.esl.CallerPrettyfier != nil {
			var fileLineVal string
			funcVal, fileLineVal = h.esl.CallerPrettyfier(entry.Caller)
			if sep := strings.IndexRune(fileLineVal, ':'); sep != -1 {
				fileVal = fileLineVal[:sep]
				lineVal, _ = strconv.Atoi(fileLineVal[sep+1:])
			} else {
				fileVal = fileLineVal
				lineVal = 0
			}
		} else {
			funcVal = entry.Caller.Function
			fileVal = entry.Caller.File
			lineVal = entry.Caller.Line
		}
		entry.Caller = nil

		if funcVal != "" {
			data["log.origin.function"] = funcVal
		}
		if fileVal != "" {
			data["log.origin.file.name"] = fileVal
		}
		if lineVal > 0 {
			data["log.origin.file.line"] = lineVal
		}
	}

	for k, v := range data {
		entry.Data[k] = v
	}

	return nil
}

func getTagName(name string, logName string) map[string]string {
	var envName string
	if strings.Contains(name, "dev") {
		envName = "dev"
	} else if strings.Contains(name, "staging") {
		envName = "staging"
	} else if strings.Contains(name, "prestaging") {
		envName = "prestaging"
	} else if strings.Contains(name, "prod") {
		envName = "prod"
	}

	data := make(map[string]string)
	data["debug"] = logName + "." + envName + ".debug"
	data["info"] = logName + "." + envName + ".info"
	data["warn"] = logName + "." + envName + ".warn"
	data["error"] = logName + "." + envName + ".error"
	data["fatal"] = logName + "." + envName + ".fatal"
	data["panic"] = logName + "." + envName + ".panic"

	return data
}
