package Models

type Card struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Word       string `json:"word"`
	Definition string `json:"definition"`
	List       string `json:"list"`
}
