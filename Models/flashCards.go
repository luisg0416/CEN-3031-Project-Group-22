package Models

type Card struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// need to change routes and creation functions to match this
// also need to change keys for gorm but we need to get the database running first
type CardGroup struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Cards []Card `json:"cards"`
}
