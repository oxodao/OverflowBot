package web

import "github.com/oxodao/overflow-bot/services"

func createOrUpdateUser(prv *services.Provider, user *DiscordUser) error {
	rq := `
		INSERT INTO DISCORD_USER (DISCORD_ID, DISCORD_USERNAME)
		VALUES ($1, $2)
		ON CONFLICT (DISCORD_ID) DO
		UPDATE SET DISCORD_USERNAME = $2, TOKEN = uuid_generate_v4()
		RETURNING TOKEN`

	row := prv.DB.QueryRow(rq, user.ID, user.Username)
	if row.Err() != nil {
		return row.Err()
	}

	token := ""
	err := row.Scan(&token)
	user.Token = token
	return err
}

func findUserByToken(prv *services.Provider, token string) (*DiscordUser, error) {
	row := prv.DB.QueryRowx("SELECT DISCORD_ID, DISCORD_USERNAME, TOKEN FROM DISCORD_USER WHERE TOKEN = $1", token)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user DiscordUser
	err := row.StructScan(&user)
	return &user, err
}
