package main

import (
	"os"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/urfave/cli"
)

func migrationStart() {
	if err := db_main.AutoMigrate(
		&pb.MappingORM{},
		&pb.CurrencyORM{},
	); err != nil {
		log.Errorln("Migration failed: %v", err)
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

			log.Println("Migration process begin...")
			migrationStart()
			log.Println("Migration process finished...")

			return nil
		},
	}
}
