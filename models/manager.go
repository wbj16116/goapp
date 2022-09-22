package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int64
	IsSuper  int
	Role     Role `gorm:"foreignkey:Id;association_foreignkey:RoleId"`  //设置关联  references 和  association_foreignkey  的区别 一对一，一对多
}

func (Manager) TableName() string {
	return "manager"
}

