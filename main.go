package main

import (
	"net/http"

	"collector/internal/commands"
	"collector/internal/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	var command commands.Command
	if err := command.Init(); err != nil {
		log.Fatal("error during initialization", zap.Error(err))
	}

	done := make(chan error)

	go func() {
		if err := command.StartTagCleaner(); err != nil {
			done <- err
		}
	}()
	go func() {
		if err := command.StartDurationChecker(); err != nil {
			done <- err
		}
	}()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9100", nil); err != nil {
			done <- err
		}
	}()

	err := <-done
	if err != nil {
		log.Fatal("runtime error", zap.Error(err))
	}
}
