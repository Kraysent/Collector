package main

import (
	"context"
	"fmt"
	"time"

	"collector/internal/actions"
	"collector/internal/core"
)

func main() {
	ctx := context.Background()
	repo, err := core.NewRepository()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Running needtime cleaner on last 100 completed tasks")
		n, err := actions.NeedtimeTagClean(ctx, repo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Done, cleaned 'needtime' tag from %d tasks\n", n)
		time.Sleep(1 * time.Minute)
	}

}
