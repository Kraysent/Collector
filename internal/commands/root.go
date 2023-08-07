package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"collector/internal/actions"
	"collector/internal/core"
	"collector/internal/log"
	"go.uber.org/zap"
)

type Command struct {
	ctx        context.Context
	Repository *core.Repository
}

func (c *Command) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.Background()
	}

	return c.ctx
}

func (c *Command) Init() error {
	configPath, ok := os.LookupEnv("CONFIG")
	if !ok {
		return fmt.Errorf("no config specified")
	}

	config, err := core.ParseConfig(configPath)
	if err != nil {
		return err
	}

	repo, err := core.NewRepository(config)
	if err != nil {
		return err
	}

	c.Repository = repo

	return log.InitLogger(config.Logging.StdoutPath, config.Logging.StderrPath)
}

func (c *Command) StartTagCleaner() error {
	if c.Repository.Config.TagCleaner.Disabled {
		return nil
	}

	for {
		log.Info("Running tag cleaner",
			zap.String("tag_name", c.Repository.Config.TagCleaner.TagName),
			zap.Int("limit", c.Repository.Config.TagCleaner.Limit),
		)
		n, err := actions.CleanTagFromCompletedTasks(c.Context(), c.Repository)
		if err != nil {
			return err
		}

		log.Info("Done, cleaned tag from tasks",
			zap.String("tag_name", c.Repository.Config.TagCleaner.TagName),
			zap.Int("number_of_affected_tasks", n),
			zap.Time("next_clean_time", time.Now().Add(c.Repository.Config.TagCleaner.Period)),
		)
		time.Sleep(c.Repository.Config.TagCleaner.Period)
	}
}

func (c *Command) StartDurationChecker() error {
	if c.Repository.Config.DurationChecker.Disabled {
		return nil
	}
	for {
		if err := actions.ReadCodeDurations(c.Context(), c.Repository); err != nil {
			panic(err)
		}
		time.Sleep(c.Repository.Config.DurationChecker.Period)
	}
}
