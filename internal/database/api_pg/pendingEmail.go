package apipg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func PendingDB(email models.PendingEmail) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}
	query := `
    INSERT INTO pending_emails (id, email)
    VALUES ($1, $2)`

	_, err := db.Exec(query, email.Id, email.Email)
	if err != nil {
		return "", fmt.Errorf("error inserting record into the database: %v", err)
	}

	success := "Data successfully inserted into the database"
	log.Println(success)
	return success, nil
}
