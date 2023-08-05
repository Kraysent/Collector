package wakatime

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"collector/internal/interactions/client"
)

const (
	DefaultEndpoint = "https://wakatime.com/api"
)

type Client interface {
	client.Client
	GetDurations(ctx context.Context, day, month, year int) ([]Duration, error)
}

var _ Client = (*ClientImpl)(nil)

type ClientImpl struct {
	client *http.Client
	token  string
}

func NewClient(token string) *ClientImpl {
	return &ClientImpl{
		client: http.DefaultClient,
		token:  token,
	}
}

func (c *ClientImpl) GetEndpoint() string {
	return DefaultEndpoint
}

func (c *ClientImpl) GetHTTPClient() *http.Client {
	return c.client
}

func (c *ClientImpl) GetDurations(ctx context.Context, day, month, year int) ([]Duration, error) {
	tokenB64 := base64.StdEncoding.EncodeToString([]byte(c.token))

	response, err := client.SendJSONRequest[DurationResponse](
		ctx, c, http.MethodGet, "v1/users/current/durations",
		client.WithQuery("date", fmt.Sprintf("%d-%d-%d", year, month, day)),
		client.WithHeader("Authorization", fmt.Sprintf("Basic %s", tokenB64)),
	)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
