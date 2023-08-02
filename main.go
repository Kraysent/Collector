package main

import (
	"context"

	"collector/internal/actions"
	"collector/internal/core"
)

func main() {
	ctx := context.Background()
	repo, err := core.NewRepository()
	if err != nil {
		panic(err)
	}

	err = actions.NeedtimeTagClean(ctx, repo)
	if err != nil {
		panic(err)
	}
}
