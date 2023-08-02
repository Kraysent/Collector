package core

import (
	"collector/internal/interactions"
)

type Repository struct {
	Clients *interactions.Clients
}

func NewRepository() (*Repository, error) {
	clients, err := interactions.NewClients()
	if err != nil {
		return nil, err
	}

	return &Repository{
		Clients: clients,
	}, nil
}
