package migrations

var Migrations = map[int]string{
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
