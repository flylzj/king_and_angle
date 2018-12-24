package model

type Blessing struct {
	Blessing 	string		`json:"blessing" binding:"required"`
}
