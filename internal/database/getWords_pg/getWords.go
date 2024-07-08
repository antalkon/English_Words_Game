package getwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

// GetTranslationsByRowNumbers получает строки из столбца translate по порядковым номерам строк
func GetRandomWord() (*models.GClassicWords, error) {
	db := connectDB.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	var word models.GClassicWords
	query := "SELECT word, translate FROM words ORDER BY RANDOM() LIMIT 1"
	err := db.QueryRow(query).Scan(&word.Word, &word.Translate)
	if err != nil {
		return nil, fmt.Errorf("error querying the database: %v", err)
	}

	log.Printf("Random word: %s, Translation: %s\n", word.Word, word.Translate)
	return &word, nil
}
