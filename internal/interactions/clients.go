package interactions

import (
	"fmt"
	"os"

	"collector/internal/interactions/ticktick"
)

type Clients struct {
	TickTick ticktick.Client
}

func NewClients() (*Clients, error) {
	ticktickToken, ok := os.LookupEnv("TICKTICK_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no TickTick token provided")
	}
	ticktickClient := ticktick.NewClient(
		ticktick.WithOAuthToken(ticktickToken),
	)

	return &Clients{
		TickTick: ticktickClient,
	}, nil
}
