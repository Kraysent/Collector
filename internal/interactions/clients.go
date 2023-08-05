package interactions

import (
	"collector/internal/interactions/ticktick"
)

type Clients struct {
	TickTick ticktick.Client
}

func NewClients(ticktickToken string) (*Clients, error) {
	ticktickClient := ticktick.NewClient(
		ticktick.WithOAuthToken(ticktickToken),
	)

	return &Clients{
		TickTick: ticktickClient,
	}, nil
}
