package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"collector/internal/actions"
	"collector/internal/core"
	"collector/internal/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	ticktickToken, ok := os.LookupEnv("TICKTICK_TOKEN")
	if !ok {
		panic("no TickTick token provided")
	}

	config, err := core.ParseConfig("configs/config.yaml")
	if err != nil {
		panic(err)
	}
	config.Clients.TickTick.Token = ticktickToken

	repo, err := core.NewRepository(config)
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
			log.Info("Running tag cleaner",
				zap.String("tag_name", repo.Config.TagCleaner.TagName),
				zap.Int("limit", repo.Config.TagCleaner.Limit),
			)
			n, err := actions.CleanTagFromCompletedTasks(ctx, repo)
			if err != nil {
				done <- err
				return
			}

			log.Info("Done, cleaned tag from tasks",
				zap.String("tag_name", repo.Config.TagCleaner.TagName),
				zap.Int("number_of_affected_tasks", n),
				zap.Time("next_clean_time", time.Now().Add(repo.Config.TagCleaner.Period)),
			)
			time.Sleep(repo.Config.TagCleaner.Period)
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
