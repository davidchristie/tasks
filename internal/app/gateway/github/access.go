package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type accessResponse struct {
	AccessToken string `json:"access_token"`
}

var ErrCouldNotCreateAccessRequest = errors.New("could not create access request")
var ErrCouldNotSendAccessRequest = errors.New("could not send access request")
var ErrCouldNotParseAccessResponse = errors.New("could not parse access response")

var httpClient = http.Client{}

func requestAccessToken(logger *log.Logger, clientID, clientSecret, code string) (string, error) {
	// Create access token request
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		logger.Printf("could not create access request: %v\n", err)
		return "", ErrCouldNotCreateAccessRequest
	}
	req.Header.Set("accept", "application/json")

	// Send access token request
	res, err := httpClient.Do(req)
	if err != nil {
		logger.Printf("could not send access request: %v\n", err)
		return "", ErrCouldNotSendAccessRequest
	}
	defer res.Body.Close()

	// Parse access token response
	var body accessResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		logger.Printf("could not parse access response: %v\n", err)
		return "", ErrCouldNotParseAccessResponse
	}

	return body.AccessToken, nil
}
