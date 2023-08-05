package actions

import (
	"context"
	"time"

	"collector/internal/core"
	"collector/internal/interactions/ticktick"
	"golang.org/x/exp/slices"
)

func CleanTagFromCompletedTasks(ctx context.Context, repo *core.Repository) (int, error) {
	start := time.Now()
	tasksToUpdate := make([]ticktick.Task, 0)

	defer func() {
		end := time.Since(start)
		repo.Metrics.TagCleanDuration.Set(float64(end.Milliseconds()))
		repo.Metrics.CleanedTasksNumber.Add(float64(len(tasksToUpdate)))
	}()

	tasks, err := repo.Clients.TickTick.GetCompletedTasks(ctx)
	if err != nil {
		return 0, err
	}

	for _, task := range tasks {
		if slices.Contains(task.Tags, repo.Config.TagCleaner.TagName) {
			tags := make([]string, 0)

			for _, tag := range task.Tags {
				if tag != repo.Config.TagCleaner.TagName {
					tags = append(tags, tag)
				}
			}

			task.Tags = tags
			tasksToUpdate = append(tasksToUpdate, task)
		}
	}

	return len(tasksToUpdate), repo.Clients.TickTick.UpdateTasks(ctx, ticktick.UpdateTaskRequest{
		Update: tasksToUpdate,
	}, repo.Config.TagCleaner.Limit)
}
