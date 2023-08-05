package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Option func(r *http.Request) error

func WithHeader(header, value string) Option {
	return func(r *http.Request) error {
		r.Header.Add(header, value)
		return nil
	}
}

func WithCookie(name, value string) Option {
	return func(r *http.Request) error {
		r.AddCookie(&http.Cookie{
			Name: name, Value: value,
		})
		return nil
	}
}

func WithQuery(key, value string) Option {
	return func(r *http.Request) error {
		q := r.URL.Query()
		q.Add(key, value)
		r.URL.RawQuery = q.Encode()
		return nil
	}
}

func WithJSONBody(data any) Option {
	return func(r *http.Request) error {
		if r.Method == http.MethodGet {
			return fmt.Errorf("body is not allowed for GET request")
		}

		bodyBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}

		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		return nil
	}
}

func WithJSONContentType() Option {
	return func(r *http.Request) error {
		return WithHeader("Content-Type", "application/json")(r)
	}
}
