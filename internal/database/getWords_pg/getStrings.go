package getwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func GetStringCount() (int, error) {
	db := connectDB.GetDB()
	if db == nil {
		return 0, errors.New("failed to connect to the database")
	}

	var count int
	query := "SELECT COUNT(*) FROM words"
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error querying the database: %v", err)
	}

	successMessage := fmt.Sprintf("There are %d rows in the words table", count)
	log.Println(successMessage)
	return count, nil
}
