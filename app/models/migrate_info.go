package models

import "time"

type MigrateInfo struct {
	Num       int       `db:"num"`
	Status    int       `db:"status"`
	EventDate time.Time `db:"event_date"`
}
