package web

type DiscordUser struct {
	ID       string `json:"id" db:"discord_id"`
	Username string `json:"username" db:"discord_username"`
	Token    string `json:"token" db:"token"`
}
