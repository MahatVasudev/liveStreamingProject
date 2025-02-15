package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MahatVasudev/liveStreamingProject/server/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.DBEnvs.DBUser,
			config.DBEnvs.DBPassword,
			config.DBEnvs.DBHost,
			config.DBEnvs.Port,
			config.DBEnvs.DBName,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd1 := os.Args[(len(os.Args) - 1)]
	cmd2 := os.Args[(len(os.Args) - 2)]
	if cmd1 == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
	if cmd1 == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
	if cmd2 == "force" {
		inn, err := strconv.ParseInt(cmd1, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if err := m.Force(int(inn)); err != nil {
			log.Fatal(err)
		}
	}
}
