package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Age      int    `json:"age" gorm:"not null"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
