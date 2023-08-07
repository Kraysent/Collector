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

	configPath, ok := os.LookupEnv("CONFIG")
	if !ok {
		panic("no config specified")
	}

	config, err := core.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	repo, err := core.NewRepository(config)
	if err != nil {
		panic(err)
	}

	if err := log.InitLogger(config.Logging.StdoutPath, config.Logging.StderrPath); err != nil {
		panic(err)
	}

	done := make(chan error)

	go func() {
		if repo.Config.TagCleaner.Disabled {
			return
		}
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
		if repo.Config.DurationChecker.Disabled {
			return
		}
		for {
			if err := actions.ReadCodeDurations(ctx, repo); err != nil {
				panic(err)
			}
			time.Sleep(repo.Config.DurationChecker.Period)
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
