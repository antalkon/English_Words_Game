package addwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func CreateWordSet(createWord models.AddSetTxt) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	// Вставка в таблицу WordSet
	insertQuery := `
    INSERT INTO WordSet (SetId, SetName, SetAuthor, code)
    VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(insertQuery, createWord.SetId, createWord.SetName, createWord.Email, createWord.Code)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error inserting record into the database: %v", err)
	}

	// Создание таблицы с динамическим именем (ensure createWord.SetId is a valid identifier)
	createTableQuery := fmt.Sprintf(`
    CREATE TABLE %s (
        WordId VARCHAR(30) PRIMARY KEY,
        Word VARCHAR(30) NOT NULL,
        Translate VARCHAR(30) NOT NULL
    )`, createWord.Code)

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return "", fmt.Errorf("error creating table in the database: %v", err)
	}

	success := "Data successfully inserted and table created in the database"
	log.Println(success)
	return success, nil
}
