package ticktick

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	DefaultEndpoint = "https://api.ticktick.com/api"
)

type Client interface {
	GetProjects(context.Context) ([]Project, error)
	GetCompletedTasks(ctx context.Context) ([]Task, error)
}

var _ Client = (*ClientImpl)(nil)

type ClientImpl struct {
	client   *http.Client
	endpoint string
	token    string
}

func NewClient(options ...Option) *ClientImpl {
	client := &ClientImpl{
		endpoint: DefaultEndpoint,
		client:   http.DefaultClient,
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

func sendAPIRequest[OutputType any](
	ctx context.Context, client *ClientImpl, method string, path string,
) (*OutputType, error) {
	request, err := http.NewRequestWithContext(
		ctx, method, strings.Join([]string{client.endpoint, path}, "/"), nil,
	)
	if err != nil {
		return nil, err
	}

	request.AddCookie(&http.Cookie{
		Name: "t", Value: client.token,
	})

	response, err := client.client.Do(request)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result OutputType
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *ClientImpl) GetProjects(ctx context.Context) ([]Project, error) {
	projects, err := sendAPIRequest[[]Project](ctx, c, http.MethodGet, "v2/projects")

	return *projects, err
}

func (c *ClientImpl) GetCompletedTasks(
	ctx context.Context,
) ([]Task, error) {
	tasks, err := sendAPIRequest[[]Task](ctx, c, http.MethodGet, "v2/project/all/completed")

	return *tasks, err
}
