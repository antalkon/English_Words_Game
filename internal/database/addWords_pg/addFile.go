package addwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func AddFileWords(id, word, translate string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}
	query := `
    INSERT INTO words (WordID, Word, Translate)
    VALUES ($1, $2, $3)`

	_, err := db.Exec(query, id, word, translate)
	if err != nil {
		return "", fmt.Errorf("error inserting record into the database: %v", err)
	}

	success := "Data successfully inserted into the database"
	log.Println(success)
	return success, nil
}
