package oauthpg

import (
	"errors"

	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func CheckDBAccByID(userID string) (string, error) {
	db := connectDB.GetDB() // Предполагается, что connectDB.GetDB() возвращает *sql.DB
	if db == nil {
		return "", errors.New("Fail connect DB")
	}

	query := "SELECT COUNT(*) FROM accounts WHERE userId = $1"
	var count int
	err := db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return "", errors.New("Fail exec from DB")
	}

	if count > 0 {
		return "true", nil
	}
	return "false", nil
}
