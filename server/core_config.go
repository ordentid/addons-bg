package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Listen address is an array of IP addresses and port combinations.
	// Listen address is an array so that this service can listen to many interfaces at once.
	// You can use this value for example: []string{"192.168.1.12:80", "25.49.25.73:80"} to listen to
	// listen to interfaces with IP address of 192.168.1.12 and 25.49.25.73, both on port 80.
	ListenAddress string `config:"LISTEN_ADDRESS"`

	CorsAllowedHeaders []string `config:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods []string `config:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS"`

	JWTSecret   string `config:"JWT_SECRET"`
	JWTDuration string `config:"JWT_DURATION"`
	Dsn         string `config:"DB_DSN"`

	TaskService     string `config:"TASK_SERVICE"`
	AuthService     string `config:"AUTH_SERVICE"`
	UserService     string `config:"USER_SERVICE"`
	CompanyService  string `config:"COMPANY_SERVICE"`
	WorkflowService string `config:"WORKFLOW_SERVICE"`
}

var config *Config

func initConfig() {
	// Todo: add env checker

	godotenv.Load(".env")
	// if err != nil {
	// 	log.Println(err)
	// 	log.Fatalf("Error loading .env file")
	// }

	config = &Config{
		ListenAddress: fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		CorsAllowedHeaders: []string{
			"Connection", "User-Agent", "Referer",
			"Accept", "Accept-Language", "Content-Type",
			"Content-Language", "Content-Disposition", "Origin",
			"Content-Length", "Authorization", "ResponseType",
			"X-Requested-With", "X-Forwarded-For",
		},
		CorsAllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		CorsAllowedOrigins: []string{"*"},
		JWTSecret:          getEnv("JWT_SECRET", "secret"),
		JWTDuration:        getEnv("JWT_DURATION", "48h"),
		Dsn:                getEnv("DB_DSN", ""),
		TaskService:        getEnv("TASK_SERVICE", ":9090"),
	}

}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
