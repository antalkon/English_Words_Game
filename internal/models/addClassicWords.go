package models

type AddClassicWords struct {
	WordId    string `json:"WordsID`
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type AddSetTxt struct {
	SetName string `json:"SetName"`
	Email   string `json:"Email"`
	SetId   string `json:"SetId"`
	Code    string `json:"Code"`
}
