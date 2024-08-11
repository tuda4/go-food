package main

import (
	"context"
	"go-food/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
)

var interrupSignal = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Info("error: ", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), interrupSignal...)
	defer stop()

	pool, err := pgxpool.New(ctx, config.DbSource)
	if err != nil {
		slog.Info("error: ", err)
	}

	if pool == nil {
		slog.Info("Could not connect to database")
	}

	defer pool.Close()

	slog.Info("Connected to database")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":8080")
}
