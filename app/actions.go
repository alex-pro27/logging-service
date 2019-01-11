package main

import (
	"./models"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

type Actions struct {
	DB *sqlx.DB
}

func (actions Actions) AddLog(c *gin.Context) {
	var err error
	text := c.PostForm("text")
	typeLog := c.PostForm("type")
	appName := c.PostForm("app_name")

	tx := actions.DB.MustBegin()

	log := models.Log{appName, text, typeLog, time.Now(), time.Now()}
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
			"error":   true,
			"message": err.Error(),
		})
	} else {
		c.IndentedJSON(200, gin.H{
			"success": true,
		})
	}
}

func (actions Actions) GetLogs(c *gin.Context) {
	var err error
	app, isApp := c.GetQuery("app")
	typeLog, isTypeLog := c.GetQuery("type")

	_logs := sq.Select("app, text, created", "type_log").From("logging.log").Limit(100)

	if isApp {
		_logs = _logs.Where(sq.Eq{"app": app})
	}

	if isTypeLog {
		_logs = _logs.Where(sq.Eq{"type_log": typeLog})
	}

	sql, args, err := _logs.ToSql()

	fmt.Println(args)

	var logs []models.Log
	err = actions.DB.Select(&logs, sql, args...)

	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	} else {
		var res []gin.H
		for _, l := range logs {
			res = append(res, gin.H{
				"app":     l.App,
				"type":    l.Type,
				"text":    l.Text,
				"created": l.Created,
			})
		}
		c.IndentedJSON(200, res)
	}

}

func (actions Actions) Ping(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}
