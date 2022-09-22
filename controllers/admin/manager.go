package admin

import (
	"github.com/astaxie/beego"
	"goapp/models"
)

type ManagerController struct {
	BaseController
}

func (c *ManagerController) Get() {
	manager := []models.Manager{}
	models.DB.Preload("Role").Find(&manager) //联表查询
	c.Data["managerList"] = manager
	// c.Data["json"] = manager
	// c.ServeJSON()
	c.TplName = "admin/manager/index.html"
}

func (c *ManagerController) Add() {
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	beego.Info(role)
	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("参数有误", "/manager")
		return
	}
	manager := models.Manager{}
	models.DB.Where("Id = ?", id).Find(&manager)

	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.Data["manager"] = manager
	c.TplName = "admin/manager/edit.html"
}

func (c *ManagerController) DoAdd() {
	username := models.Trim(c.GetString("username"))
	password := models.Trim(c.GetString("password"))
	roleId, error   := c.GetInt("role_id")
	if error != nil {
		c.Error("角色ID有误", "/manager/add")
		return
	}
	mobile   := models.Trim(c.GetString("mobile"))
	email    := models.Trim(c.GetString("email"))

	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或密码不合法", "/manager/add")
		return
	}
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile: mobile,
		Email:email,
		RoleId:roleId,
		Status:1,
		AddTime:models.GetIntTime(),
	}
	err := models.DB.Create(&manager).Error
	if err == nil {
		c.Success("添加管理员成功", "/mananger")
	} else {
		c.Error("添加角管理员败", "/mananger/add")
	}
}

func (c *ManagerController) DoEdit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("ID有误", "/manager/edit")
		return
	}
	username := models.Trim(c.GetString("username"))
	password := models.Trim(c.GetString("password"))
	roleId, error   := c.GetInt("role_id")
	if error != nil {
		c.Error("角色ID有误", "/manager/add")
		return
	}
	mobile   := models.Trim(c.GetString("mobile"))
	email    := models.Trim(c.GetString("email"))

	if len(username) < 2{
		c.Error("用户名", "/manager/add")
		return
	}
	manager := models.Manager{Id:id}
	models.DB.Find(&manager)
	manager.Username = username
	if password != "" {
		if len(password) < 6 {
			c.Error("密码不合法", "/manager/edit?id=" + models.IntToString(id))
			return
		}
		manager.Password = models.Md5(password)
	}
	manager.Mobile = mobile
	manager.Email = email
	manager.RoleId = roleId
	err := models.DB.Save(&manager).Error
	if err == nil {
		c.Success("编辑管理员成功", "/manager")
	} else {
		c.Error("添加管理员失败", "/manager/edit?id=" + models.IntToString(id))
	}
}

func (c *ManagerController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("ID有误", "/manager/edit")
		return
	}
	manager := models.Manager{Id:id}
	err := models.DB.Delete(&manager).Error
	if err == nil {
		c.Success("删除管理员成功", "/manager")
	} else {
		c.Error("删除管理员失败", "/manager")
	}
}
