package main

import (
	"os"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func migrationStart() {
	if err := db_main.AutoMigrate(
		&pb.TransactionORM{},
	); err != nil {
		logrus.Fatalf("Migration failed: %v", err)
		os.Exit(1)
	}
}

func runMigrationCmd() cli.Command {
	return cli.Command{
		Name:  "db-migrate",
		Usage: "run db migration",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			initDBMain()
			defer closeDBMain()

			logrus.Println("Migration process begin...")
			migrationStart()
			logrus.Println("Migration process finished...")

			return nil
		},
	}
}
