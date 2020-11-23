package web

import (
	"net/http"

	"github.com/oxodao/overflow-bot/services"
)

func call(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://discordapp.com/api/oauth2/authorize?scope=bot&client_id="+prv.Config.Discord.ClientID, http.StatusTemporaryRedirect)
	}
}
