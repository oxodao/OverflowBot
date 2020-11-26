package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oxodao/overflow-bot/log"
	"github.com/oxodao/overflow-bot/models"
	"github.com/oxodao/overflow-bot/services"
)

func FetchCommand(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmds, err := listCommands(prv)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, _ := json.Marshal(&cmds)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}
}

func CreateCommand(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(5 << 20)

		name := r.Form.Get("name")
		help := r.Form.Get("help")
		resp := r.Form.Get("resp")

		if len(name) == 0 || len(help) == 0 || len(resp) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Trying to create a command with no name/help/response")
			return
		}

		row := prv.DB.QueryRowx(`INSERT INTO CUSTOM_COMMANDS(COMMAND_NAME, COMMAND_HELP, COMMAND_RESPONSE) VALUES ($1, $2, $3) RETURNING COMMAND_ID, COMMAND_NAME, COMMAND_HELP, COMMAND_RESPONSE`, name, help, resp)
		if row.Err() != nil {
			// @TODO: check for conflict and tell the user
			w.WriteHeader(http.StatusBadRequest)
			log.Error(row.Err())
			return
		}

		var cmd models.Command
		err := row.StructScan(&cmd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		bytes, err := json.Marshal(&cmd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		prv.ReloadCommands <- true

		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}

func UpdateCommand(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := r.ParseMultipartForm(2 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		name := r.Form.Get("name")
		help := r.Form.Get("help")
		resp := r.Form.Get("resp")

		res, err := prv.DB.Exec(`
			UPDATE CUSTOM_COMMANDS
			SET 
				COMMAND_NAME = CASE WHEN LENGTH($2) > 0 THEN $2 ELSE COMMAND_NAME END,
				COMMAND_HELP = CASE WHEN LENGTH($3) > 0 THEN $3 ELSE COMMAND_HELP END,
				COMMAND_RESPONSE = CASE WHEN LENGTH($4) > 0 THEN $4 ELSE COMMAND_RESPONSE END
			WHERE COMMAND_ID = $1`, id, name, help, resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		ra, err := res.RowsAffected()
		if err != nil || ra == 0 {
			w.WriteHeader(http.StatusBadRequest)
			if ra == 0 {
				log.Error("No rows affected")
			} else {
				log.Error(err)
			}
			return
		}

		idInt, _ := strconv.Atoi(id)
		c, err := getCommand(prv, idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		prv.ReloadCommands <- true

		str, _ := json.Marshal(c)
		w.Write([]byte(str))
	}
}

func DeleteCommand(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		_, err := prv.DB.Exec("DELETE FROM CUSTOM_COMMANDS WHERE COMMAND_ID = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't remove the database entry!"))
			log.Error(err)
			return
		}

		prv.ReloadCommands <- true

		w.WriteHeader(http.StatusOK)
	}
}
