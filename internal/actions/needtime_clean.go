package actions

import (
	"context"

	"collector/internal/core"
	"collector/internal/interactions/ticktick"
	"golang.org/x/exp/slices"
)

const (
	NeedtimeTag = "needtime"
)

func NeedtimeTagClean(ctx context.Context, repo *core.Repository) error {
	tasks, err := repo.Clients.TickTick.GetCompletedTasks(ctx)
	if err != nil {
		return err
	}

	tasksToUpdate := make([]ticktick.Task, 0)

	for _, task := range tasks {
		if slices.Contains(task.Tags, NeedtimeTag) {
			tags := make([]string, 0)

			for _, tag := range task.Tags {
				if tag != NeedtimeTag {
					tags = append(tags, tag)
				}
			}

			task.Tags = tags
			tasksToUpdate = append(tasksToUpdate, task)
		}
	}

	return repo.Clients.TickTick.UpdateTasks(ctx, ticktick.UpdateTaskRequest{
		Update: tasksToUpdate,
	})
}
