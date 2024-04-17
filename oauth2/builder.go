package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Builder is a structure used to configure middleware
type Builder struct {
	jwksURL        string
	issuer         string
	debug          bool
	audience       string
	httpClient     HttpClient
	requestMatcher map[string]map[string][]string
	allowUnmatched bool
	clientAudience string
	clientId       string
	clientSecret   string
	tokenUrl       string
}

// Option type for the configuring middleware builder
type Option func(*Builder)

// Config prepares builder
func (o *Builder) Config(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// URL stores the JWKS URL in the builder
func URL(jwksURL string) Option {
	return func(auth2 *Builder) {
		auth2.jwksURL = jwksURL
	}
}

// HTTPClient stores the http client in the middleware
func HTTPClient(httpClient HttpClient) Option {
	return func(auth2 *Builder) {
		auth2.httpClient = httpClient
	}
}

// Debug enables debug logging
func Debug(debug bool) Option {
	return func(auth2 *Builder) {
		auth2.debug = debug
	}
}

// Unmatched configure middleware to allow unmatched method/path combinations
// to access respective resource
func Unmatched(allowUnmatched bool) Option {
	return func(auth2 *Builder) {
		auth2.allowUnmatched = allowUnmatched
	}
}

// Issuer stores the issuer expected in the tokens
func Issuer(issuer string) Option {
	return func(auth2 *Builder) {
		auth2.issuer = issuer
	}
}

// Audience stores the audience expected in the tokens
func Audience(audience string) Option {
	return func(auth2 *Builder) {
		auth2.audience = audience
	}
}

// Request stores a new request matcher for the method, path and optional set of scopes
func Request(method, path string, scopes []string) Option {
	return func(auth2 *Builder) {
		if auth2.requestMatcher == nil {
			auth2.requestMatcher = make(map[string]map[string][]string)
		}

		pathRequestMatcher := auth2.requestMatcher[path]
		if pathRequestMatcher == nil {
			pathRequestMatcher = make(map[string][]string)
			auth2.requestMatcher[path] = pathRequestMatcher
		}

		pathRequestMatcher[method] = scopes
	}
}

// Build creates middleware configuration structure
func Build(opts ...Option) (*Config, error) {
	builder := &Builder{
		debug: false,
	}

	builder.Config(opts...)

	ctx := context.Background()
	if builder.httpClient == nil {
		return nil, errors.New("http client is required property, use WithHttpClient builder function to set it up")
	}

	if builder.jwksURL == "" {
		return nil, errors.New("JWKS URL is not set, cannot continue")
	}

	config := &Config{
		debug:          builder.debug,
		issuer:         builder.issuer,
		audience:       builder.audience,
		requestMatcher: builder.requestMatcher,
		allowUnmatched: builder.allowUnmatched,
	}

	request, err := http.NewRequestWithContext(ctx, "GET", builder.jwksURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create ID provider: %v", err)
	}

	response, err := builder.httpClient.Do(request)
	if err != nil || response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get content from the ID provider: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %v", err)
	}

	err = json.Unmarshal(body, &config.jksConfiguration)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response body: %v", err)
	}
	log.Println("the OAuth2 middleware has been successfully initialized")

	return config, nil
}
