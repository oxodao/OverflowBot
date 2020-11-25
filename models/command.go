package models

type Command struct {
	ID   int    `json:"id" db:"command_id"`
	Name string `json:"name" db:"command_name"`
	Help string `json:"help" db:"command_help"`
	Resp string `json:"resp" db:"command_response"`
}
