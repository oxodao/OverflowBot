package models

type Sound struct {
	ID          int    `json:"id" db:"sound_id"`
	Name        string `json:"name" db:"sound_name"`
	File        string `json:"file" db:"sound_file"`
	Description string `json:"desc" db:"sound_desc"`
}
