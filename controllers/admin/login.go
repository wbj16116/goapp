package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"goapp/models"
	"strings"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
	//models.DB.LogMode(true)  //开启sql日志
}

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	beego.Info(c.GetSession("userinfo"))
	c.TplName = "admin/login/login.html"
}

func (c *LoginController) DoLogin() {
	var flag = cpt.VerifyReq(c.Ctx.Request)

	if flag {
		username := strings.Trim(c.GetString("username"), "")
		password := strings.Trim(models.Md5(c.GetString("password")), "")
		beego.Info(username)
		manager := []models.Manager{}
		models.DB.Where("username=? AND password=?", username, password).Find(&manager)
		if (len(manager)>0) {
			c.SetSession("userinfo", manager[0])
			beego.Info(c.GetSession("userinfo"))
			c.Success("登录成功", "")
		} else {
			c.Success("用户名或密码错误", "/login")
		}

	} else {
		c.Error("验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出成功", "/login")
}
