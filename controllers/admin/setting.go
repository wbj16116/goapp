package admin

import (
	"goapp/models"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Get() {
	setting := models.Setting{}
	models.DB.First(&setting)       //First
	c.Data["setting"] = setting
	c.TplName = "admin/setting/index.html"
}

func (c *SettingController) DoEdit() {
	//1、获取数据库里面的数据
	setting := models.Setting{}
	models.DB.Find(&setting)
	//2、修改数据
	c.ParseForm(&setting)   //把表单提交的字段，直接转换到结构体中。不用一个一个下定义。依托  SiteTitle       string `form:"site_title"`
	//上传图片 site_logo
	siteLogo, err1 := c.UploadImg("site_logo")
	if len(siteLogo) > 0 && err1 == nil {
		setting.SiteLogo = siteLogo
	}
	//上传图片no_picture
	noPicture, err2 := c.UploadImg("no_picture")
	if len(noPicture) > 0 && err2 == nil {
		setting.NoPicture = noPicture
	}
	//执行保存数据
	err3 := models.DB.Where("id=1").Save(&setting).Error
	if err3 != nil {
		c.Error("修改数据失败", "/setting")
		return
	}
	c.Success("修改数据成功", "/setting")

}
