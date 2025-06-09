package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Course    string `json:"course"`
	Track     string `json:"track"`
	Status    string `json:"status"`
}
