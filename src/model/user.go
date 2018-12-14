package model

type User struct {
	ID			uint		`grom:"AUTO_INCREMENT"`
	Username 	string		`gorm:"type:varchar(10)"`
	Password 	string
	Name		string
	Sex       	uint
	KingUsername	string
	Wish 		string		`gorm:"type:text"`
	Blessing    string		`gorm:"type:text"`
}

func (User) TableName()string{
	return "user"
}
