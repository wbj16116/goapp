package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"goapp/models"
	"strings"
	"net/url"
)

func AdminAuth(ctx *context.Context) {
	pathname := ctx.Request.URL.String()
	userinfo, ok := ctx.Input.Session("userinfo").(models.Manager)
	if !(ok && userinfo.Username != "") {
		if pathname != beego.AppConfig.String("adminPath")+"/login" && pathname != beego.AppConfig.String("adminPath")+"/login/doLogin" {
			ctx.Redirect(302, beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		pathname := strings.Replace(pathname, beego.AppConfig.String("adminPath"), "", 1)  //获取url和参数
		urlPath,_ := url.Parse(pathname) //去掉url后的参数

		if userinfo.IsSuper == 0 && !excludeAuthPath(string(urlPath.Path)) {
			roleId := userinfo.RoleId
			roleAccess := []models.RoleAccess{}
			models.DB.Where("role_id=?", roleId).Find(&roleAccess)
			roleAccessMap := make(map[int]int)
			for _, v := range roleAccess {
				roleAccessMap[v.AccessId] = v.AccessId
			}

			//获取当前url对应的权限ID
			access := models.Access{}
			models.DB.Where("url=?", urlPath.Path).Find(&access)

			//判断当前url对应的权限ID是否在当前角色的权限ID中
			if _, ok := roleAccessMap[access.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}

//判断当前链接是否需要验证权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(beego.AppConfig.String("excludeAuthPath"), ",") //字符串切割为数组，等于php中的explode()
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return false
		}
	}
	return false
}
