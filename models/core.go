package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	DB, _ = gorm.Open("mysql", "wbj:123456@tcp(120.27.24.29:3306)/mybeego?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Error()
	}
}
