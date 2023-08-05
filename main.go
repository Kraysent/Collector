package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"collector/internal/actions"
	"collector/internal/core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()
	repo, err := core.NewRepository()
	if err != nil {
		panic(err)
	}

	done := make(chan error)

	go func() {
		for {
			fmt.Println("Running needtime cleaner on last 100 completed tasks")
			n, err := actions.NeedtimeTagClean(ctx, repo)
			if err != nil {
				done <- err
				return
			}

			fmt.Printf("Done, cleaned 'needtime' tag from %d tasks\n", n)
			time.Sleep(1 * time.Minute)
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
