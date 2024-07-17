package oauthpg

import (
	"errors"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func AddUser(userData models.OauthUser) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("fail connect DB")
	}

	query := "INSERT INTO accounts (userId, email, phone, Name, Surname) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, userData.ID, userData.Email, userData.Phone, userData.Name, userData.Surname)
	if err != nil {
		return "", errors.New("fail exec DB")
	}

	return "success", nil
}
