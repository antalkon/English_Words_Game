package models

type GClassicWords struct {
	Word      string
	Translate string
}

type GClassicSetW struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type GClassicSet struct {
	Set string `json:"set"`
}
