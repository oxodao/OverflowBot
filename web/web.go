package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/oxodao/overflow-bot/services"
)

func Initialize(prv *services.Provider) {
	r := mux.NewRouter()

	r.HandleFunc("/call", call(prv))

	srv := &http.Server{
		Addr:         "0.0.0.0:" + prv.Config.Web.Port,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
