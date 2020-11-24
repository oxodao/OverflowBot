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

	// Actual api
	api.HandleFunc("/cours", FetchCoursHandler(prv))
	api.HandleFunc("/sounds", FetchSounds(prv))

	api.HandleFunc("/sound", CreateSound(prv)).Methods(http.MethodPost)
	api.HandleFunc("/sound/{id}", HearSound(prv)).Methods(http.MethodGet) // Not used yet, too lazy to it in VueJS (requires a service worker to add the token to the request + needs to forward the token to the service worker)
	api.HandleFunc("/sound/{id}", UpdateSound(prv)).Methods(http.MethodPut)
	api.HandleFunc("/sound/{id}", DeleteSound(prv)).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + prv.Config.Web.Port,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
