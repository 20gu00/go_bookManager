package model

type User struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Token    string  `json:"token"`
	Books    []*Book `gorm:"many2many:book_user"` //写一个就够了,自动生成多对多关联表,也可以自定义多对多关联表
}

func (*User) TableName() string {
	return "user"
}
