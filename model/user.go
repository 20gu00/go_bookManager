package model

type User struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Token    string  `json:"token"`
	Books    []*Book `gorm:"many2many:book_users"`
}

func (*User) TableName() string {
	return "user"
}
