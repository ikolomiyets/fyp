package auth0

import (
	"errors"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Builder is a structure used to configure middleware
type Builder struct {
	baseUrl      string
	clientId     string
	clientSecret string
	debug        bool
	audience     string
	httpClient   HttpClient
}

// Option type for the configuring middleware builder
type Option func(*Builder)

// Config prepares builder
func (o *Builder) Config(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// Audience stores the audience expected in the tokens
func Audience(audience string) Option {
	return func(auth0 *Builder) {
		auth0.audience = audience
	}
}

// ClientId stores the client secret for the Auth0 management API client
func ClientId(clientID string) Option {
	return func(auth0 *Builder) {
		auth0.clientId = clientID
	}
}

// ClientSecret stores the client secret for the Auth0 management API client
func ClientSecret(clientSecret string) Option {
	return func(auth0 *Builder) {
		auth0.clientSecret = clientSecret
	}
}

// BaseUrl stores the token URL for obtaining the client access token
func BaseUrl(baseUrl string) Option {
	return func(auth0 *Builder) {
		auth0.baseUrl = baseUrl
	}
}

// HTTPClient stores the http client in the middleware
func HTTPClient(httpClient HttpClient) Option {
	return func(auth0 *Builder) {
		auth0.httpClient = httpClient
	}
}

// Debug set the debug flag
func Debug(debug bool) Option {
	return func(auth0 *Builder) {
		auth0.debug = debug
	}
}

// Build creates Auth0 client structure
func Build(opts ...Option) (*Config, error) {
	builder := &Builder{
		debug: false,
	}

	builder.Config(opts...)

	if builder.httpClient == nil {
		return nil, errors.New("http client is required property, use HttpClient builder function to set it up")
	}

	config := &Config{
		baseUrl:      builder.baseUrl,
		debug:        builder.debug,
		audience:     builder.audience,
		clientId:     builder.clientId,
		clientSecret: builder.clientSecret,
		httpClient:   builder.httpClient,
	}

	return config, nil
}
