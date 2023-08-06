package actions

import (
	"context"
	"time"

	"collector/internal/core"
	"collector/internal/storage"
)

func ReadCodeDurations(ctx context.Context, repo *core.Repository) error {
	currTime := time.Now()
	durations, err := repo.Clients.WakaTime.GetDurations(ctx, currTime.Day(), int(currTime.Month()), currTime.Year())
	if err != nil {
		return err
	}

	events := make([]storage.Event, 0)

	for _, dur := range durations {
		events = append(events, storage.Event{
			Timestamp: time.Unix(int64(dur.Timestamp), 0),
			Source:    "wakatime",
			Data: map[string]any{
				"project":  dur.Project,
				"duration": dur.Duration,
			},
		})
	}

	return repo.Storage.EventStorage.InsertEvents(ctx, events)
}
