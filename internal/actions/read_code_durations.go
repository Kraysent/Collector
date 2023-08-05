package actions

import (
	"context"
	"time"

	"collector/internal/core"
	"collector/internal/log"
	"go.uber.org/zap"
)

func ReadCodeDurations(ctx context.Context, repo *core.Repository) error {
	currTime := time.Now()
	durations, err := repo.Clients.WakaTime.GetDurations(ctx, currTime.Day(), int(currTime.Month()), currTime.Year())
	if err != nil {
		return err
	}

	for _, dur := range durations {
		log.Info("Read duration",
			zap.Time("timestamp", time.Unix(int64(dur.Timestamp), 0)),
			zap.String("project", dur.Project),
			zap.Duration("duration", time.Duration(dur.Duration*float64(time.Second))),
		)
	}

	return nil
}
