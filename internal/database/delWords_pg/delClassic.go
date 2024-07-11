package delwordspg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func DelClassicWords(delWord models.DelClassicWord) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	// Запрос для удаления строки
	deleteQuery := `
    DELETE FROM words
    WHERE Word = $1`

	// Выполнение запроса на удаление
	_, err := db.Exec(deleteQuery, delWord.Word)
	if err != nil {
		return "", fmt.Errorf("error deleting record from the database: %v", err)
	}

	success := "Data successfully deleted from the database"
	log.Println(success)
	return success, nil
}
