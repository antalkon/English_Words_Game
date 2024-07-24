package datapg

import (
	"errors"

	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func CheckUserByID(id string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("Connect fail")
	}

	query := "SELECT id FROM accounts WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return "", errors.New("NF")
	}
	return "Yes", nil
}
