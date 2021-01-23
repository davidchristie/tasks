package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/davidchristie/tasks/internal/app/gateway/config"
)

type User struct {
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

func FetchUser(conf *config.Config, token string) (*User, error) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, conf.GithubUserURL, nil)
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
