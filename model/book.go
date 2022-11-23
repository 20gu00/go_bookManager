package model

type Book struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Desc  string  `json:"desc"`
	Users []*User `gorm:"many2many:book_user"`
}

func (*Book) TabelName() string {
	return "book"
}
