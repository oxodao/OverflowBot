package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/oxodao/overflow-bot/config"
)

type Provider struct {
	Config   *config.Config
	Software *software
	DB       *sqlx.DB
	Discord  *discordgo.Session
}

type software struct {
	Version string
	Author  string
}

func NewProvider(cfg *config.Config, version, author string) (*Provider, error) {
	db, err := sqlx.Connect("postgres", cfg.DB)
	if err != nil {
		return nil, err
	}
	fmt.Println("\t- Database loaded")

	return &Provider{
		Config: cfg,
		Software: &software{
			Author:  author,
			Version: version,
		},
		DB: db,
	}, nil
}
