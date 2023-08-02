package ticktick

type Option func(client *ClientImpl)

func WithEndpoint(endpoint string) Option {
	return func(client *ClientImpl) {
		client.endpoint = endpoint
	}
}

func WithOAuthToken(token string) Option {
	return func(client *ClientImpl) {
		client.token = token
	}
}
