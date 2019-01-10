package main

import "time"

type migrateInfo struct {
	Num int `db:"num"`
	Status int `db:"status"`
	EventDate time.Time `db:"event_date"`
}

var migrations = map[int] string{
	0: `
		CREATE TABLE IF NOT EXISTS logging.log (
			app String,
			text Text,
    		type_log String,
    		created DateTime,
    		event_date Date
		) ENGINE=MergeTree(event_date, (app, created), 8192)
	`,
}