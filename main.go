package main

import (
	"context"
	"net/http"
	"time"

	"collector/internal/actions"
	"collector/internal/core"
	"collector/internal/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	repo, err := core.NewRepository()
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log.SetLogger(*logger)

	done := make(chan error)

	go func() {
		for {
			log.Info("Running needtime cleaner on last 100 completed tasks")
			n, err := actions.NeedtimeTagClean(ctx, repo)
			if err != nil {
				done <- err
				return
			}

			log.Info("Done, cleaned 'needtime' tag from tasks", zap.Int("number_of_tasks", n))
			time.Sleep(1 * time.Minute)
		}
	}()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9100", nil); err != nil {
			done <- err
		}
	}()

	err = <-done
	if err != nil {
		panic(err)
	}
}
