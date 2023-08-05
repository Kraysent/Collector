package core

import (
	"collector/internal/interactions"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Repository struct {
	Clients *interactions.Clients
	Metrics struct {
		NeedtimeCleanDuration prometheus.Gauge
	}
}

func NewRepository() (*Repository, error) {
	clients, err := interactions.NewClients()
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		Clients: clients,
	}

	repo.Metrics.NeedtimeCleanDuration = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "needtime_clean_duration",
			Help: "Duration of the cleaning needtime tag process",
		},
	)

	return repo, nil
}
