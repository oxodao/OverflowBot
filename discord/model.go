package discord

import "time"

type Cours struct {
	Name string    `db:"cours_name"`
	Date time.Time `db:"cours_date"`
	End  time.Time `db:"cours_end"`
}

type Sound struct {
	Name        string `db:"sound_name"`
	File        string `db:"sound_file"`
	Description string `db:"sound_desc"`
}
