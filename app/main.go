package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	fmt.Println("CLICKHOUSE_SERVER", os.Getenv("CLICKHOUSE_SERVER"))

	db, err := sqlx.Open(
		"clickhouse",
		os.Getenv("CLICKHOUSE_SERVER"),
	)

	handlerError(err)

	actions := Actions{
		DB: db,
	}

	api := router.Group("/api/")
	{
		api.GET("ping/", actions.Ping)
		api.PUT("add-log/", actions.AddLog)
		api.GET("logs/", actions.GetLogs)
	}

	handlerError(router.Run(":8080"))
}

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
