package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseMigrations string `required:"true" split_words:"true"`
	DatabaseURL string `required:"true" split_words:"true"`
	Domain string `required:"true"`
	Environment string `default:"production"`
	GithubAccessTokenURL string `default:"https://github.com/login/oauth/access_token" split_words:"true"`
	GithubAuthorizeURL string `default:"https://github.com/login/oauth/authorize" split_words:"true"`
	GithubClientID string `required:"true" split_words:"true"`
	GithubClientSecret string `required:"true" split_words:"true"`
	GithubUserURL string `default:"https://api.github.com/user" split_words:"true"`
	Port int `required:"true"`
	WebApp string `required:"true" split_words:"true"`
}

func Read() *Config {
	conf := &Config{}
	envconfig.MustProcess("", conf)
	return conf
}
