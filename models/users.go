package models

type Users struct {
	Id       int64
	Username string `orm:"size(128)"`
	Phone    int64
	Name     string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	Password string `orm:"size(128)"`
}

func (Users) TableName() string {
	return "users"
}

