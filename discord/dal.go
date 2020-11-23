package discord

import "github.com/oxodao/overflow-bot/services"

func SelectCustomCommands(prv *services.Provider) ([]Command, error) {
	cmds := []Command{}
	rows, err := prv.DB.Queryx("SELECT COMMAND_NAME, COMMAND_HELP, COMMAND_RESPONSE FROM CUSTOM_COMMANDS ORDER BY COMMAND_NAME")
	if err != nil {
		return cmds, err
	}

	for rows.Next() {
		c := CustomCommand{}
		rows.StructScan(&c)

		cmds = append(cmds, c)
	}

	return cmds, nil
}

func SelectCours(prv *services.Provider) ([]Cours, error) {
	cours := []Cours{}

	rows, err := prv.DB.Queryx(`
	(
		SELECT COURS_NAME, COURS_DATE, COURS_END
		FROM COURS
		WHERE COURS_DATE > NOW() 
		ORDER BY COURS_DATE
		LIMIT 1
	)
	UNION
	(
		SELECT COURS_NAME, COURS_DATE, COURS_END
		FROM COURS
		WHERE COURS_DATE < NOW()
		ORDER BY COURS_DATE DESC
		LIMIT 1
	)
	ORDER BY COURS_DATE`)

	if err != nil {
		return cours, err
	}

	for rows.Next() {
		c := Cours{}
		rows.StructScan(&c)
		cours = append(cours, c)
	}

	return cours, nil
}

func SelectSounds(prv *services.Provider) ([]Sound, error) {
	rows, err := prv.DB.Queryx("SELECT SOUND_NAME, SOUND_FILE, SOUND_DESC FROM SOUNDS ORDER BY SOUND_NAME")
	if err != nil {
		return []Sound{}, err
	}

	var sounds []Sound
	for rows.Next() {
		s := Sound{}
		rows.StructScan(&s)
		sounds = append(sounds, s)
	}

	return sounds, nil
}

func SelectSound(prv *services.Provider, name string) (*Sound, error) {
	row := prv.DB.QueryRowx("SELECT SOUND_NAME, SOUND_FILE, SOUND_DESC FROM SOUNDS WHERE LOWER(SOUND_NAME) = $1", name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var sound Sound
	err := row.StructScan(&sound)
	return &sound, err
}
