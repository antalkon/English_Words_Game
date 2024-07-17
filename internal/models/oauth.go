package models

type OauthUser struct {
	ID      string `json:"userId"`
	Email   string `json:"userEmail"`
	Phone   string `json:"userPhone"`
	Name    string `json:"userName"`
	Surname string `json:"userSurname"`
	Exp     int64  `json:"exp"`
}
