package models

type Card struct {
	Id         int    `json:"id"`
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// need to change routes and creation functions to match this
// also need to change keys for gorm but we need to get the database running first
type CardWUser struct {
	Id         int    `csv:"id"`
	User       string `csv:"user"`
	Word       string `csv:"word"`
	Definition string `csv:"definition"`
}