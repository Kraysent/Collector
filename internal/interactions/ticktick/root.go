package ticktick

import (
	"context"
	"net/http"
	"strconv"

	"collector/internal/interactions/client"
)

const (
	DefaultEndpoint = "https://api.ticktick.com/api"
)

type Client interface {
	client.Client
	GetProjects(context.Context) ([]Project, error)
	GetCompletedTasks(context.Context) ([]Task, error)
	UpdateTasks(context.Context, UpdateTaskRequest, int) error
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

func (c *ClientImpl) getCommonOptions() []client.Option {
	return []client.Option{
		client.WithCookie("t", c.token),
		client.WithJSONContentType(),
	}
}

func (c *ClientImpl) GetProjects(ctx context.Context) ([]Project, error) {
	projects, err := client.SendJSONRequest[[]Project](
		ctx, c, http.MethodGet, "v2/projects",
		c.getCommonOptions()...,
	)
	if err != nil {
		return nil, err
	}

	return *projects, nil
}

func (c *ClientImpl) GetCompletedTasks(ctx context.Context) ([]Task, error) {
	tasks, err := client.SendJSONRequest[[]Task](
		ctx, c, http.MethodGet, "v2/project/all/completed",
		c.getCommonOptions()...,
	)

	if err != nil {
		return nil, err
	}

	return *tasks, nil
}

func (c *ClientImpl) UpdateTasks(ctx context.Context, updateRequest UpdateTaskRequest, limit int) error {
	_, err := client.SendJSONRequest[any](
		ctx, c, http.MethodPost, "v2/batch/task",
		append(c.getCommonOptions(),
			client.WithJSONBody(updateRequest),
			client.WithQuery("limit", strconv.Itoa(limit)),
		)...,
	)

	return err
}
