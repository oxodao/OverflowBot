package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/oxodao/overflow-bot/services"
)

func FetchCours(prv *services.Provider, fileToLoad string) error {
	f, err := os.Open(fileToLoad)
	if err != nil {
		return err
	}

	prv.DB.Exec("DROP FROM cours")

	r := csv.NewReader(f)
	firstLine := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
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
		//fmt.Printf("Cours: %v @ %v jusque %v\n", nom, start, end)
	}

	return nil
}

func ToRealDate(dateCnam string) string {
	dates := strings.Split(dateCnam, "/")
	return dates[2] + "-" + dates[1] + "-" + dates[0]
}
