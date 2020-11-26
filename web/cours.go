package web

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"strings"

	"github.com/oxodao/overflow-bot/log"
	"github.com/oxodao/overflow-bot/services"
)

func FetchCours(prv *services.Provider, buf bytes.Buffer) error {
	prv.DB.Exec("DROP FROM cours")

	reader := strings.NewReader(buf.String())
	r := csv.NewReader(reader)
	firstLine := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error(err)
			return err
		}

		if firstLine {
			firstLine = false
			continue
		}

		nom := record[16]
		start := ToRealDate(record[1]) + " " + record[2]
		end := ToRealDate(record[3]) + " " + record[4]

		prv.DB.Exec(`INSERT INTO COURS(COURS_NAME, COURS_DATE, COURS_END) VALUES ($1, $2, $3)`, nom, start, end)
	}

	return nil
}

func FetchCoursHandler(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)

		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		defer file.Close()

		var buf bytes.Buffer
		io.Copy(&buf, file)

		err = FetchCours(prv, buf)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func ToRealDate(dateCnam string) string {
	dates := strings.Split(dateCnam, "/")
	return dates[2] + "-" + dates[1] + "-" + dates[0]
}
