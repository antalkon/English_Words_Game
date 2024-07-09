package getwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func GetRandomSet(set string) (*models.GClassicSetW, error) {
	db := connectDB.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	var word models.GClassicSetW
	query := fmt.Sprintf(`SELECT word, translate FROM %s ORDER BY RANDOM() LIMIT 1`, set)
	err := db.QueryRow(query).Scan(&word.Word, &word.Translate)
	if err != nil {
		return nil, fmt.Errorf("error querying the database: %v", err)
	}

	log.Printf("Random word: %s, Translation: %s\n", word.Word, word.Translate)
	return &word, nil
}
