package main

import (
	"../migrations"
	"../models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"log"
	"os"
	"time"
)

func main() {
	command := os.Args[1]

	switch command {
	case "migrate":
		migrate()
		break
	default:
		log.Fatal(fmt.Sprintf("command %s not defined, use one of: {migrate}", command))
	}
}

func Database() (*sqlx.DB, error) {
	return sqlx.Open(
		"clickhouse",
		os.Getenv("CLICKHOUSE_SERVER")+"?debug=true",
	)
}

func migrate() {
	var err error
	db, err := Database()
	if err != nil {
		log.Fatal(err)
	}

	var _migratesInfo []models.MigrateInfo
	err = db.Select(&_migratesInfo, `SELECT * FROM migrations`)

	migratesInfo := make(map[int]int)
	for _, migrateInfo := range _migratesInfo {
		migratesInfo[migrateInfo.Num] = migrateInfo.Status
	}

	tx := db.MustBegin()
	i := 0
	for num, schema := range migrations.Migrations {
		status, ok := migratesInfo[num]
		if !ok || status != 1 {
			fmt.Printf("migrate %d\n", num)
			db.MustExec(schema)
			m := models.MigrateInfo{Num: num, Status: 1, EventDate: time.Now()}
			_, err = tx.NamedExec(
				`
				INSERT INTO migrations
				(num, status, event_date) 
				VALUES (:num, :status, :event_date)
				`,
				m,
			)
			handlerError(err)
			handlerError(tx.Commit())
			fmt.Println("Ok")
			i++
		}
	}
	if i == 0 {
		fmt.Println("No migrate apply")
	}
}

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
