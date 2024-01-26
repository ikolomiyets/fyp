package oauth2

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

// Config contains middleware configuration
type Config struct {
	jksConfiguration Jwks
	issuer           string
	debug            bool
	audience         string
	requestMatcher   map[string]map[string][]string
	allowUnmatched   bool
}

func (o *Config) parseToken(ctx context.Context, tokenString string) (string, error) {
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
		return "", errors.New("invalid token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		log.Println("token has invalid signature")
		return "", errors.New("invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired):
		log.Println("token has expired")
		return "", errors.New("token has expired")
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		if nb, ok := token.Claims.GetNotBefore(); ok == nil {
			if nb.After(time.Now().Add(-30 * time.Second)) {
				return extractClaims(ctx, token)
			}
		}
		log.Println("token is not valid yet")
		return "", errors.New("token is not valid yet")
	default:
		log.Println("error parsing JWT", "error", err)
		return "", err
	}
}

func extractClaims(ctx context.Context, token *jwt.Token) (string, error) {
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		var scope string
		if scope, ok = mapClaims["scope"].(string); ok {
			return scope, nil
		}
		// Cater for Microsoft token for the time being
		if scope, ok = mapClaims["scp"].(string); ok {
			return scope, nil
		}
		return "", nil
	}
	return "", errors.New("not map claims")
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

func (o *Config) validateScopes(ctx context.Context, authorizationHeader string, scopes []string) (bool, error) {
	if authorizationHeader == "" && scopes == nil {
		return true, nil
	}

	tokenString, err := extractToken(ctx, authorizationHeader)
	if err != nil {
		return false, err
	}

	if tokenString == "" {
		return false, nil
	}

	scope, err := o.parseToken(ctx, tokenString)
	if err != nil {
		return false, err
	}

	// If scopes are not provided authenticated user should be granted access
	if scopes == nil || len(scopes) == 0 {
		return true, nil
	}

	if scope != "" {
		result := strings.Split(scope, " ")
		for i := range result {
			for j := range scopes {
				if result[i] == scopes[j] {
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
