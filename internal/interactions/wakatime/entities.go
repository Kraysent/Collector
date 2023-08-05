package wakatime

type Duration struct {
	Timestamp float64 `json:"time"`
	Project   string  `json:"project"`
	Duration  float64 `json:"duration"`
}

type DurationResponse struct {
	Data []Duration `json:"data"`
}
