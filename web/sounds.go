package web

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oxodao/overflow-bot/log"
	"github.com/oxodao/overflow-bot/models"
	"github.com/oxodao/overflow-bot/services"
)

func FetchSounds(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sounds, err := listSounds(prv)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		bytes, _ := json.Marshal(&sounds)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}
}

func CreateSound(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(5 << 20)

		name := r.Form.Get("name")
		desc := r.Form.Get("desc")
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		if len(name) == 0 || len(desc) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		row := prv.DB.QueryRowx(`INSERT INTO SOUNDS(SOUND_NAME, SOUND_FILE, SOUND_DESC) VALUES ($1, $2, $3) RETURNING SOUND_ID, SOUND_NAME, SOUND_FILE, SOUND_DESC`, name, name+".mp3", desc)
		if row.Err() != nil {
			// @TODO: check for conflict and tell the user
			w.WriteHeader(http.StatusBadRequest)
			log.Error(row.Err())
			return
		}

		var snd models.Sound
		err = row.StructScan(&snd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		dest, err := os.Create("sounds/" + snd.File)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't create file"))
			log.Error(err)
			return
		}
		defer dest.Close()

		_, err = io.Copy(dest, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		bytes, err := json.Marshal(&snd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}

func HearSound(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		idInt, _ := strconv.Atoi(id)

		s, err := getSound(prv, idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		http.ServeFile(w, r, "sounds/"+s.File)
	}
}

func UpdateSound(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := r.ParseMultipartForm(2 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		name := r.Form.Get("name")
		desc := r.Form.Get("desc")

		res, err := prv.DB.Exec(`
			UPDATE SOUNDS
			SET 
				SOUND_NAME = CASE WHEN LENGTH($2) > 0 THEN $2 ELSE SOUND_NAME END,
				SOUND_DESC = CASE WHEN LENGTH($3) > 0 THEN $3 ELSE SOUND_NAME END
			WHERE SOUND_ID = $1`, id, name, desc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		ra, err := res.RowsAffected()
		if err != nil || ra == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		idInt, _ := strconv.Atoi(id)
		s, err := getSound(prv, idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		str, _ := json.Marshal(s)
		w.Write([]byte(str))
	}
}

func DeleteSound(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		row := prv.DB.QueryRowx("SELECT SOUND_ID, SOUND_NAME, SOUND_FILE, SOUND_DESC FROM SOUNDS WHERE SOUND_ID = $1", id)
		if row.Err() != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Error(row.Err())
			return
		}

		var sound models.Sound
		row.StructScan(&sound)

		err := os.Remove("sounds/" + sound.File)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't remove the file!"))
			log.Error(err)
			return
		}

		_, err = prv.DB.Exec("DELETE FROM SOUNDS WHERE SOUND_ID = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't remove the database entry!"))
			log.Error(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
