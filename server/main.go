package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"google.golang.org/grpc"

	"fmt"
	"net"

	svc "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/spf13/viper"
)

const defaultPort = 9090
const serviceName = "BG"

var s *grpc.Server

func main() {

	initConfig()

	app := cli.NewApp()
	app.Name = ""
	app.Commands = []cli.Command{
		grpcServerCmd(),
		gatewayServerCmd(),
		grpcGatewayServerCmd(),
		runMigrationCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func grpcServerCmd() cli.Command {
	return cli.Command{
		Name:  "grpc-server",
		Usage: "starts a gRPC server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: defaultPort,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")

			startDBConnection()

			go func() {
				if err := grpcServer(port); err != nil {
					logrus.Fatalf("failed RPC serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			closeDBConnections()

			logrus.Println("Stopping RPC server")
			s.Stop()
			logrus.Println("RPC server stopped")
			return nil
		},
	}
}

func gatewayServerCmd() cli.Command {
	return cli.Command{
		Name:  "gw-server",
		Usage: "starts a Gateway server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 3000,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: ":" + fmt.Sprint(defaultPort),
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			port, grpcEndpoint := c.Int("port"), c.String("grpc-endpoint")

			go func() {
				if err := httpGatewayServer(port, grpcEndpoint); err != nil {
					logrus.Fatalf("failed JSON Gateway serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			logrus.Println("JSON Gateway server stopped")

			return nil
		},
	}
}

func grpcGatewayServerCmd() cli.Command {
	return cli.Command{
		Name:  "grpc-gw-server",
		Usage: "Starts gRPC and Gateway server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port1",
				Value: defaultPort,
			},
			cli.IntFlag{
				Name:  "port2",
				Value: 3000,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: ":" + fmt.Sprint(defaultPort),
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			rpcPort, httpPort, grpcEndpoint := c.Int("port1"), c.Int("port2"), c.String("grpc-endpoint")

			startDBConnection()

			go func() {
				if err := grpcServer(rpcPort); err != nil {
					logrus.Fatalf("failed RPC serve: %v", err)
				}
			}()

			go func() {
				if err := httpGatewayServer(httpPort, grpcEndpoint); err != nil {
					logrus.Fatalf("failed JSON Gateway serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			logrus.Println("Stopping RPC server")
			s.GracefulStop()
			closeDBConnections()
			logrus.Println("RPC server stopped")
			logrus.Println("JSON Gateway server stopped")

			return nil
		},
	}
}

func grpcServer(port int) error {
	// RPC
	logrus.Printf("Starting %s Service ................", serviceName)
	logrus.Printf("Starting RPC server on port %d...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// logrus.Println("===========> %s", taskConn.GetState().String())

	svcConn := svc.InitServicesConn(
		"",
		config.TaskService,
		config.AuthService,
		config.CompanyService,
		config.WorkflowService,
		config.SystemService,
		config.TransactionService,
		config.AccountService,
		config.MenuService,
		config.UserService,
		config.CutOffService,
		config.NotificationService,
	)
	defer svcConn.CloseAllServicesConn()

	apiServer := api.New(
		config.JWTSecret,
		config.JWTDuration,
		db_main,
		svcConn,
	)
	authInterceptor := api.NewAuthInterceptor(apiServer.GetManager())

	unaryInterceptorOpt := grpc.UnaryInterceptor(api.UnaryInterceptors(authInterceptor))
	streamInterceptorOpt := grpc.StreamInterceptor(api.StreamInterceptors(authInterceptor))

	s = grpc.NewServer(unaryInterceptorOpt, streamInterceptorOpt)
	pb.RegisterApiServiceServer(s, apiServer)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s.Serve(list)
}

func httpGatewayServer(port int, grpcEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	rmux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomHTTPError),
		runtime.WithForwardResponseOption(httpResponseModifier),
	)
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := pb.RegisterBaseServiceHandlerFromEndpoint(ctx, rmux, grpcEndpoint, opts)
	client := pb.NewApiServiceClient(conn)
	err = pb.RegisterApiServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		return err
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", originMiddleware(rmux))

	mux.HandleFunc("/api/bg/docs/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/api/bg/docs/", http.StripPrefix("/api/bg/docs/", fs))

	// Start
	logrus.Printf("Starting JSON Gateway server on port %d...", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), cors(mux))
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}

func originMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		referer := r.Header.Get("Referer")
		envOrigin := r.Header.Get("ENV-Allow-Origin")
		envOrigins := strings.Split(envOrigin, ",")
		for i, v := range envOrigins {
			envOrigins[i] = strings.TrimSpace(v)
		}

		logrus.Infof("Origin: %v - Ref: %v - ENV: %v", origin, referer, envOrigins)

		if getEnv("ENV", "DEV") == "PROD" {
			pass := false
			if origin != "" {
				for _, v := range envOrigins {
					if origin == v {
						pass = true
					}
				}
			}
			if referer != "" {
				for _, v := range envOrigins {
					if strings.Contains(referer, v) {
						pass = true
					}
				}
			}
			if !pass {
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string) bool {
	if stringInSlice(viper.GetString("cors"), config.CorsAllowedOrigins) {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		if allowedOrigin(r.Header.Get("Origin")) {
			if getEnv("ENV", "DEV") != "PROD" {
				w.Header().Set("Content-Security-Policy", "object-src 'none'; child-src 'none'; script-src 'unsafe-inline' https: http: ")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.Header().Set("X-Frame-Options", "DENY")
				w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
				w.Header().Set("X-XSS-Protection", "1; mode=block")
				w.Header().Set("Permissions-Policy", "geolocation=()")
				w.Header().Set("Referrer-Policy", "no-referrer")

				w.Header().Set("Access-Control-Allow-Origin", strings.Join(config.CorsAllowedOrigins, ", "))
			}
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.CorsAllowedMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.CorsAllowedHeaders, ", "))
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
