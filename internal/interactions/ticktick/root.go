package ticktick

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"collector/internal/log"
	"go.uber.org/zap"
)

const (
	DefaultEndpoint = "https://api.ticktick.com/api"
)

type Client interface {
	GetProjects(context.Context) ([]Project, error)
	GetCompletedTasks(context.Context) ([]Task, error)
	UpdateTasks(context.Context, UpdateTaskRequest, int) error
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
	ctx context.Context, client *ClientImpl, method string, path string, query map[string]string, body any,
) (*OutputType, error) {
	var bodyReader io.Reader

	if method != http.MethodGet {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(bodyBytes)
	}

	target := strings.Join([]string{client.endpoint, path}, "/")

	request, err := http.NewRequestWithContext(
		ctx, method, target, bodyReader,
	)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	for key, val := range query {
		q.Add(key, val)
	}
	request.URL.RawQuery = q.Encode()

	if bodyReader != nil {
		request.Header.Add("Content-Type", "application/json")
	}

	request.AddCookie(&http.Cookie{
		Name: "t", Value: client.token,
	})

	log.Debug("Sending request",
		zap.String("target", target), zap.String("method", method), zap.Any("query", query),
	)

	response, err := client.client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Debug("Got response", zap.ByteString("response_body", responseBytes))

	var result OutputType
	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *ClientImpl) GetProjects(ctx context.Context) ([]Project, error) {
	projects, err := sendAPIRequest[[]Project](
		ctx, c, http.MethodGet, "v2/projects", nil, nil,
	)
	if err != nil {
		return nil, err
	}

	return *projects, nil
}

func (c *ClientImpl) GetCompletedTasks(ctx context.Context) ([]Task, error) {
	tasks, err := sendAPIRequest[[]Task](
		ctx, c, http.MethodGet, "v2/project/all/completed", nil, nil,
	)
	if err != nil {
		return nil, err
	}

	return *tasks, nil
}

func (c *ClientImpl) UpdateTasks(ctx context.Context, updateRequest UpdateTaskRequest, limit int) error {
	_, err := sendAPIRequest[any](
		ctx, c, http.MethodPost, "v2/batch/task", map[string]string{
			"limit": strconv.Itoa(limit),
		}, updateRequest,
	)

	return err
}
