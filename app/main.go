package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"os"
	"time"
)

func main() {
	runServer()
}

func runServer() {
	router := gin.Default()
	fmt.Println("CLICKHOUSE_SERVER", os.Getenv("CLICKHOUSE_SERVER"))
	api := router.Group("/api/")
	{
		api.GET("ping/", ping)
		api.PUT("add-log/", addLog)
		api.GET("logs/", getLogs)
	}
	_ = router.Run(":8080")
}

type Log struct {
	App string `db:"app" json:"app"`
	Text string `db:"text" json:"text"`
	Type string `db:"type_log" json:"type"`
	Created time.Time `db:"created" json:"created"`
	EventDate time.Time `db:"event_date" json:"event_date"`
}

func Database() (*sqlx.DB, error) {
	return sqlx.Open(
		"clickhouse",
		os.Getenv("CLICKHOUSE_SERVER") + "?debug=true",
	)
}

func addLog(c *gin.Context) {
	var err error
	text := c.PostForm("text")
	typeLog := c.PostForm("type")
	appName := c.PostForm("app_name")

	db, err := Database()
	tx := db.MustBegin()

	log := Log{appName, text, typeLog, time.Now(), time.Now()}
	_, err = tx.NamedExec(
		`
		INSERT INTO logging.log
		(app, text, type_log, created, event_date) 
		VALUES (:app, :text, :type_log, :created, :event_date)
		`,
		&log,
	)
	err = tx.Commit()

	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
			"message": err.Error(),
		})
	} else {
		c.IndentedJSON(200, gin.H{
			"message": "ok",
		})
	}
}

func getLogs(c *gin.Context) {
	var err error

	db, err := Database()

	var logs []Log
	err = db.Select(&logs, `SELECT app, text, created FROM logging.log LIMIT 100`)

	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": true,
			"message": err.Error(),
		})
	} else {
		var res []gin.H
		for _, l := range logs {
			res = append(res, gin.H{
				"app": l.App,
				"text": l.Text,
				"created": l.Created,
			})
		}
		c.IndentedJSON(200, res)
	}

}


func ping(c *gin.Context)  {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}