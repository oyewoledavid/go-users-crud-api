package models

type User struct {
	UserID    string   `json:"user_id" gorm:"primaryKey;size:8"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Course    string `json:"course"`
	Track     string `json:"track"`
	Username  string `json:"username" gorm:"unique"`
	Account   Account `json:"account" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;references:UserID"`
}

type Account struct {
	UserID   string `json:"user_id" gorm:"uniqueIndex;size:8"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}
