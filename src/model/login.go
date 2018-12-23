package model

type LoginModel struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}

type PasswordModel struct {
	Password	string		`json:"password" binding:"required"`
}