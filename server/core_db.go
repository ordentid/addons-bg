package main

import (
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB or Collection
var (
	db_main     *gorm.DB
	db_main_sql *sql.DB
)

func startDBConnection() {
	logrus.Printf("Starting Db Connections...")

	initDBMain()

}

func closeDBConnections() {
	closeDBMain()
}

func initDBMain() {
	logrus.Printf("Main Db - Connecting")
	var err error
	db_main, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed connect to DB main: %v", err)
		os.Exit(1)
		return
	}

	db_main_sql, err = db_main.DB()
	if err != nil {
		logrus.Fatalf("Error cannot initiate connection to DB main: %v", err)
		os.Exit(1)
		return
	}

	db_main_sql.SetMaxIdleConns(0)
	db_main_sql.SetMaxOpenConns(100)

	err = db_main_sql.Ping()
	if err != nil {
		logrus.Fatalf("Cannot ping DB main: %v", err)
		os.Exit(1)
		return
	}

	logrus.Printf("Main Db - Connected")
}

func closeDBMain() {
	logrus.Print("Closing DB Main Connection ... ")
	if err := db_main_sql.Close(); err != nil {
		logrus.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	logrus.Println("Closing DB Main Success")
}
