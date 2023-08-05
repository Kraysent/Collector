package interactions

import (
	"collector/internal/interactions/ticktick"
	"collector/internal/interactions/wakatime"
)

type Clients struct {
	TickTick ticktick.Client
	WakaTime wakatime.Client
}

func NewClients(ticktickToken string, wakatimeToken string) (*Clients, error) {
	ticktickClient := ticktick.NewClient(ticktickToken)
	wakatimeClient := wakatime.NewClient(wakatimeToken)

	return &Clients{
		TickTick: ticktickClient,
		WakaTime: wakatimeClient,
	}, nil
}
