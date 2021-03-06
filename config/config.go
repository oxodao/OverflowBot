package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	DB      string        `json:"db"`
	Discord DiscordConfig `json:"discord"`
	Web     WebConfig     `json:"web"`
}

type DiscordConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`

	Token string `json:"token"`
}

type WebConfig struct {
	Url          string   `json:"url"`
	Port         string   `json:"port"`
	AllowedUsers []string `json:"allowedUsers"`
}

func (c *Config) Save() error {
	cfg, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("config.json", cfg, os.ModePerm)
}

func Load() (*Config, error) {
	str, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(str, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Web.Url[len(cfg.Web.Url)-1:] != "/" {
		cfg.Web.Url = cfg.Web.Url + "/"
	}

	return &cfg, nil
}
