package oauth2

import (
	"FYP/security"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"log/slog"
	"strings"
	"time"
)

// Config contains middleware configuration
type Config struct {
	jksConfiguration     Jwks
	issuer               string
	debug                bool
	audience             string
	requestMatcher       map[string]map[string][]string
	allowUnmatched       bool
	clientAudience       string
	clientId             string
	clientSecret         string
	tokenUrl             string
	accessToken          string
	accessTokenExpiresAt time.Time
}

func (o *Config) parseToken(ctx context.Context, tokenString string) (*security.Authority, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		cert, err := o.GetCert(ctx, token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	switch {
	case token.Valid:
		return extractClaims(ctx, token)
	case errors.Is(err, jwt.ErrTokenMalformed):
		log.Println("this is not a valid token")
		return nil, errors.New("invalid token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		log.Println("token has invalid signature")
		return nil, errors.New("invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired):
		log.Println("token has expired")
		return nil, errors.New("token has expired")
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		if nb, ok := token.Claims.GetNotBefore(); ok == nil {
			if nb.After(time.Now().Add(-30 * time.Second)) {
				return extractClaims(ctx, token)
			}
		}
		log.Println("token is not valid yet")
		return nil, errors.New("token is not valid yet")
	default:
		log.Println("error parsing JWT", "error", err)
		return nil, err
	}
}

func extractClaims(ctx context.Context, token *jwt.Token) (*security.Authority, error) {
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		var (
			authority security.Authority
			scopes, r string
			roles     []any
		)
		if authority.UserID, ok = mapClaims["sub"].(string); !ok {
			slog.Error("cannot extract user ID")
			return nil, errors.New("cannot extract user id")
		}

		if roles, ok = mapClaims["https://fyp.com/roles"].([]any); ok {
			for _, role := range roles {
				if r, ok = role.(string); ok {
					authority.Roles = append(authority.Roles, r)
				}
			}
		}

		if scopes, ok = mapClaims["scope"].(string); ok {
			result := strings.Split(scopes, " ")
			for _, scope := range result {
				authority.Scopes = append(authority.Scopes, scope)
			}
		}
		// Cater for Microsoft token for the time being
		if scopes, ok = mapClaims["scp"].(string); ok {
			result := strings.Split(scopes, " ")
			for _, scope := range result {
				authority.Scopes = append(authority.Scopes, scope)
			}
		}
		return &authority, nil
	}
	return nil, errors.New("not map claims")
}

// GetCert gets a certificate from the jwt.Token
func (o *Config) GetCert(ctx context.Context, token *jwt.Token) (string, error) {
	cert := ""

	for k := range o.jksConfiguration.Keys {
		if token.Header["kid"] == o.jksConfiguration.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + o.jksConfiguration.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func (o *Config) validateScopes(ctx context.Context, authority *security.Authority, scopes []string) (bool, error) {

	if authority == nil && scopes == nil {
		return true, nil
	}

	if len(authority.Scopes) > 0 {
		for i := range authority.Scopes {
			for j := range scopes {
				if authority.Scopes[i] == scopes[j] {
					return true, nil
				}
			}
		}
	} else {
		return true, nil
	}

	return false, nil
}

func extractToken(ctx context.Context, authorizationHeader string) (string, error) {
	if authorizationHeader == "" {
		return "", ErrUnauthorizedRequest
	}

	if strings.HasPrefix(authorizationHeader, "Bearer") {
		authHeaderParts := strings.Split(authorizationHeader, " ")

		if authHeaderParts[0] != "Bearer" {
			return "", nil
		}

		return authHeaderParts[1], nil
	} else {
		return authorizationHeader, nil
	}
}
