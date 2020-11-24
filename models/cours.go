package models

import "time"

type Cours struct {
	Name string    `db:"cours_name"`
	Date time.Time `db:"cours_date"`
	End  time.Time `db:"cours_end"`
}
