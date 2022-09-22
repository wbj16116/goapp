package main

import (
	"goapp/models"
	_ "goapp/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

//因为 session 内部采用了 gob 来注册存储的对象，例如 struct，所以如果你采用了非 memory 的引擎，请自己在 main.go 的 init 里面注册需要保存的这些结构体，不然会引起应用重启之后出现无法解析的错误
// 错误：gob: name not registered for interface: "/models.Manager"
func init() {
	gob.Register(models.Manager{})
}

func main() {
	//配置session存入redis
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	gob.Register(models.Manager{})
	//注册模板函数  (可在模板中直接通过管道方法调用)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("UnixToDate", models.UnixToDate)
	beego.AddFuncMap("setting", models.GetSettingFromColumn)

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)  //配置日志文件
	beego.Run()
	defer models.DB.Close()
}
