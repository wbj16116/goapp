package routers

import (
	"goapp/controllers/itying"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &itying.IndexController{})
	beego.Router("/user", &itying.UserController{})
	beego.Router("/category_:id([0-9]+).html", &itying.ProductController{},"get:CategoryList")
	beego.Router("/product", &itying.ProductController{})
}
