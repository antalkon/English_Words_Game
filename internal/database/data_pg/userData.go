package datapg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func UserData(id string) (*models.UserData, error) {
	db := connectDB.GetDB()
	if db == nil {
		return nil, errors.New("ошибка подключения к базе данных")
	}

	query := "SELECT email, phone, name, surname FROM accounts WHERE userId = $1"
	row := db.QueryRow(query, id)

	var user models.UserData
	err := row.Scan(&user.Email, &user.Phone, &user.Name, &user.Surname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("нет данных для указанного id")
		}
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}

	return &user, nil
}
