package auth0

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

// Config contains middleware configuration
type Config struct {
	baseUrl              string
	debug                bool
	audience             string
	clientId             string
	clientSecret         string
	accessToken          string
	httpClient           HttpClient
	accessTokenExpiresAt time.Time
}

// UserCreateRequest contains details of the new user request
type UserCreateRequest struct {
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	VerifyEmail   bool    `json:"verify_email"` // false
	FirstName     *string `json:"given_name,omitempty"`
	LastName      *string `json:"family_name,omitempty"`
	Name          string  `json:"name"`
	Connection    string  `json:"connection"`     // Username-Password-Authentication
	EmailVerified bool    `json:"email_verified"` // true
}

// AddRoleRequest contains details of the new role request
type AddRoleRequest struct {
	Roles []string `json:"roles"`
}

type UserQueryResponse struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
}

type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (c *Config) retrieveAccessToken() (*token, error) {
	options := url.Values{}
	options.Add("audience", c.audience)
	options.Add("client_id", c.clientId)
	options.Add("client_secret", c.clientSecret)
	options.Add("grant_type", "client_credentials")

	slog.Debug("requesting token from the ID provider", "client_id", c.clientId, "client_secret", c.clientSecret, "audience", c.audience)

	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"/oauth/token", bytes.NewBuffer([]byte(options.Encode())))
	if err != nil {
		return nil, errors.New("failed to build new POST request")
	}

	if c.debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			slog.Debug("error dumping request", "error", err)
		} else {
			slog.Debug(string(dump))
		}
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(options.Encode())))

	response, err := c.httpClient.Do(request)
	defer func() {
		_ = response.Body.Close()
	}()

	if err != nil {
		return nil, errors.New("failed to execute POST request")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code retured from Auth0: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("cannot read response body")
	}

	var content token
	err = json.Unmarshal(body, &content)
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot unmarshal response body"))
	}
	return &content, nil
}

func (c *Config) getAccessToken() (string, error) {
	if c.accessToken != "" && c.accessTokenExpiresAt.After(time.Now()) {
		return c.accessToken, nil
	}

	t, err := c.retrieveAccessToken()
	if err != nil {
		return "", errors.Join(err, errors.New("cannot retrieve access token"))
	}

	c.accessToken = t.AccessToken
	c.accessTokenExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)

	return c.accessToken, nil

}

func (c *Config) AddRole(ctx context.Context, userId string, roleId string) error {
	addRoleRequest := AddRoleRequest{
		Roles: []string{roleId},
	}

	accessToken, err := c.getAccessToken()
	if err != nil {
		return err
	}

	slog.Debug("access token", "access_token", accessToken)
	body, err := json.Marshal(addRoleRequest)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"/api/v2/users/"+userId+"/roles", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+accessToken)

	response, err := c.httpClient.Do(request)
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != 204 {
		return fmt.Errorf("cannot add role, status code %d", response.StatusCode)
	} else {
		return nil
	}
}

func (c *Config) DoesUserExist(ctx context.Context, email string) (bool, error) {
	accessToken, err := c.getAccessToken()
	if err != nil {
		return false, err
	}

	slog.Debug("access token", "access_token", accessToken)
	request, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/api/v2/users-by-email?fields=user_id%2Cemail&email="+url.QueryEscape(email), nil)
	if err != nil {
		return false, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+accessToken)

	response, err := c.httpClient.Do(request)
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("cannot read response body", "error", err)
			return false, errors.New("cannot read user info response")
		}

		slog.Debug("response", "body", string(body))

		return false, fmt.Errorf("cannot validate if user '%s' does exist", email)
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("cannot read response body", "error", err)
			return false, errors.New("cannot read user info response")
		}

		var rr []UserQueryResponse
		err = json.Unmarshal(body, &rr)
		if err != nil {
			return false, errors.New("cannot unmarshal read user info response")
		}

		if len(rr) > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (c *Config) AddUser(ctx context.Context, r UserCreateRequest) (string, error) {
	accessToken, err := c.getAccessToken()
	if err != nil {
		return "", err
	}

	slog.Debug("access token", "access_token", accessToken)
	body, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"/api/v2/users", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+accessToken)

	response, err := c.httpClient.Do(request)
	defer func() {
		_ = response.Body.Close()
	}()

	if err != nil {
		return "", fmt.Errorf("failed to create user '%s': %v", r.Email, err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != 201 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("cannot read response body", "error", err)
			return "", errors.New("cannot read create user response")
		}

		return "", fmt.Errorf("cannot add user: %s", string(body))
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("cannot read response body", "error", err)
			return "", errors.New("cannot read create user response")
		}

		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			slog.Error("cannot unmarshal response body", "error", err)
			return "", errors.New("cannot unmarshal create user response")
		}

		return jsonMap["user_id"].(string), nil
	}
}
