package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Simplyphotons/fyp.git/model"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func (c Controller) AuthorizeHandler(ctx *fiber.Ctx) error { // c Controller is selector, which means handlers can use any functions from controller

	var requestData model.AuthorizationRequest
	err := json.Unmarshal(ctx.Body(), &requestData)
	if err != nil {
		//log.Printf("cannot read request body: %v", err)
		return ctx.Status(400).JSON("cannot read request body")
	}

	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		errorMessage := model.ErrorMessage{
			Message: "Incorrect service configuration: CLIENT_ID is mandatory environment variable",
		}

		return ctx.Status(500).JSON(errorMessage)
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		errorMessage := model.ErrorMessage{
			Message: "Incorrect service configuration: CLIENT_SECRET is mandatory environment variable",
		}

		return ctx.Status(500).JSON(errorMessage)
	}

	audience := os.Getenv("AUDIENCE")
	if audience == "" {
		errorMessage := model.ErrorMessage{
			Message: "Incorrect service configuration: audience is mandatory environment variable",
		}

		return ctx.Status(500).JSON(errorMessage)
	}

	redirectUrl := os.Getenv("REDIRECT_URL")
	if redirectUrl == "" {
		errorMessage := model.ErrorMessage{
			Message: "Incorrect service configuration: redirect_url is mandatory environment variable",
		}

		return ctx.Status(500).JSON(errorMessage)
	}

	data := url.Values{} //for auth code NOT Refresh Token
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "authorization_code")
	data.Add("audience", audience)
	data.Add("code", requestData.Code)
	data.Add("redirect_uri", redirectUrl)

	httpClient := http.Client{}

	requestUrl := os.Getenv("REQUEST_URL")
	if requestUrl == "" {
		errorMessage := model.ErrorMessage{
			Message: "Incorrect service configuration: request URL is mandatory environment variable",
		}

		return ctx.Status(500).JSON(errorMessage)
	}

	request, err := http.NewRequestWithContext(ctx.Context(), "POST", requestUrl, bytes.NewBuffer([]byte(data.Encode())))
	if err != nil {
		log.Printf("cannot create token request: %v", err)
		return errors.New("cannot create token request")
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("cannot retrieve access token: %v", err)
		return errors.New("cannot retrieve access token")
	}

	if response.StatusCode != 200 {
		log.Printf("unexpected status code %d while obtaining access toke", response.StatusCode)
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("cannot read response body: %v", err)
			errorMessage := model.ErrorMessage{
				Message: "cannot read Auth0 response body",
			}
			return ctx.Status(500).JSON(errorMessage)
		}

		return ctx.Status(response.StatusCode).Send(body)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("cannot read response body: %v", err)
		return errors.New("cannot read response body")
	}

	return ctx.Status(200).Send(body)
}
