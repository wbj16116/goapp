package routers

import (
	"goapp/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/login", &api.LoginController{})
}
