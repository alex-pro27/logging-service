package models

import "time"

type Log struct {
	App       string    `db:"app" json:"app"`
	Text      string    `db:"text" json:"text"`
	Type      string    `db:"type_log" json:"type"`
	Created   time.Time `db:"created" json:"created"`
	EventDate time.Time `db:"event_date" json:"event_date"`
}
