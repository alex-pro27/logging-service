package main

import (
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
		log.Fatal(fmt.Sprintf("command %s not defined, use one of: {create-schema}", command))
	}
}

func Database() (*sqlx.DB, error) {
	return sqlx.Open(
		"clickhouse",
		os.Getenv("CLICKHOUSE_SERVER") + "?debug=true",
	)
}


func migrate() {
	db, err := Database()
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS logging.migrations (
			num Int32,
			status Int16,
			event_date Date
		) ENGINE=MergeTree(event_date, (num), 8192)
	`)

	var _migratesInfo []migrateInfo
	err = db.Select(&_migratesInfo, `SELECT * FROM logging.migrations`)

	migratesInfo := make(map[int]int)
	for _, migrateInfo := range _migratesInfo {
		migratesInfo[migrateInfo.Num] = migrateInfo.Status
	}

	tx := db.MustBegin()
	i := 0
	for num, schema := range migrations {
		status, ok := migratesInfo[num]
		if !ok || status != 1 {
			fmt.Printf("migrate %d\n", num)
			db.MustExec(schema)
			m := migrateInfo{num, 1, time.Now()}
			tx.NamedExec(
				`
				INSERT INTO logging.migrations
				(num, status, event_date) 
				VALUES (:num, :status, :event_date)
				`, &m,
			)
			tx.Commit()
			i++
			fmt.Println("Ok")
		}
	}
	if i == 0 {
		fmt.Println("No migrate apply")
	}
}




