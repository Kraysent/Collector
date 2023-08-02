package actions

import (
	"context"
	"fmt"

	"collector/internal/core"
	"collector/internal/interactions/ticktick"
	"golang.org/x/exp/slices"
)

func NeedtimeTagClean(ctx context.Context, repo *core.Repository) error {
	tasks, err := repo.Clients.TickTick.GetCompletedTasks(ctx)
	if err != nil {
		return err
	}

	eventTasks := make([]ticktick.Task, 0)

	for _, task := range tasks {
		if slices.Contains(task.Tags, "event") {
			eventTasks = append(eventTasks, task)
		}
	}

	fmt.Println(eventTasks)

	return nil
}
