package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oxodao/overflow-bot/services"
	"golang.org/x/oauth2"
)

// @TODO: Security state for OAuth2

func loginOauth(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, oauthConfig.AuthCodeURL(""), http.StatusTemporaryRedirect)
	}
}

func callbackOauth(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := oauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Step 4: Use the access token, here we use it to get the logged in user's info.
		res, err := oauthConfig.Client(oauth2.NoContext, token).Get("https://discordapp.com/api/v6/users/@me")
		if err != nil || res.StatusCode != 200 {
			w.WriteHeader(http.StatusInternalServerError)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(res.Status))
			}
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var user DiscordUser
		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing discord's response"))
			return
		}

		found := false
		for _, allowedUser := range prv.Config.Web.AllowedUsers {
			if allowedUser == user.ID {
				found = true
			}
		}

		if !found {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("This user is not allowed to used the admin panel"))
			return
		}

		err = createOrUpdateUser(prv, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't login to the server: " + err.Error()))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/#/login/%v/%v/%v", user.ID, user.Username, user.Token), http.StatusTemporaryRedirect)
	}
}

func validateToken(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*DiscordUser)
		bts, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bts)
	}
}
