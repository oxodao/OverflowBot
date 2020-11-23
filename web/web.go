package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"

	"github.com/oxodao/overflow-bot/services"
)

var oauthConfig *oauth2.Config

func Initialize(prv *services.Provider) {
	oauthConfig = &oauth2.Config{
		RedirectURL:  prv.Config.Web.Url + "auth/callback",
		ClientID:     prv.Config.Discord.ClientID,
		ClientSecret: prv.Config.Discord.ClientSecret,
		Scopes:       []string{"identify"},
		Endpoint:     Endpoint,
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(authedMiddleware(prv))
	api.HandleFunc("/call", call(prv))
	api.HandleFunc("/auth", loginOauth(prv))
	api.HandleFunc("/auth/callback", callbackOauth(prv))
	api.HandleFunc("/auth/validate", validateToken(prv))

	srv := &http.Server{
		Addr:         "0.0.0.0:" + prv.Config.Web.Port,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
