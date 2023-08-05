package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"collector/internal/log"
	"go.uber.org/zap"
)

type Client interface {
	GetEndpoint() string
	GetHTTPClient() *http.Client
}

func SendJSONRequest[ResponseBodyType any](
	ctx context.Context, client Client, method string, path string, options ...Option,
) (*ResponseBodyType, error) {
	target := strings.Join([]string{client.GetEndpoint(), path}, "/")

	request, err := http.NewRequestWithContext(
		ctx, method, target, nil,
	)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		if err := opt(request); err != nil {
			return nil, err
		}
	}

	log.Debug("Sending request",
		zap.String("target", target), zap.String("method", method),
	)

	response, err := client.GetHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Debug("Got response", zap.ByteString("response_body", responseBytes))

	var result ResponseBodyType
	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
