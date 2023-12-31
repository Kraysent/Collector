package core

import (
	"collector/internal/interactions"
	"collector/internal/storage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Repository struct {
	Clients *interactions.Clients
	Storage struct {
		EventStorage storage.EventStorage
	}
	Config  *Config
	Metrics struct {
		TagCleanDuration   prometheus.Gauge
		CleanedTasksNumber prometheus.Counter
	}
}

func NewRepository(config *Config) (*Repository, error) {
	clients, err := interactions.NewClients(config.Clients.TickTick.Token, config.Clients.WakaTime.Token)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		Clients: clients,
		Config:  config,
	}

	eventStorage, err := storage.NewStorage(config.Storage)
	if err != nil {
		return nil, err
	}
	repo.Storage.EventStorage = eventStorage

	repo.Metrics.TagCleanDuration = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tag_clean_duration",
			Help: "Duration of the cleaning needtime tag process",
		},
	)
	repo.Metrics.CleanedTasksNumber = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "number_of_cleaned_tasks",
			Help: "Number tasks which had their tags cleaned",
		},
	)

	return repo, nil
}
