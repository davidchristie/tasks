package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	AvatarURL string `json:"avatar_url"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

func FetchUser(token string) (*User, error) {
	const url = "https://api.github.com/user"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code: " + res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
