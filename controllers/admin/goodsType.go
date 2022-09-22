package admin

import (
	"goapp/models"
	"strings"
	"strconv"
	//"github.com/astaxie/beego"
)

type GoodsTypeController struct {
	BaseController
}

func (c *GoodsTypeController) Get() {
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.TplName = "admin/goodsType/index.html"
}

func (c *GoodsTypeController) Add() {
	c.TplName = "admin/goodsType/add.html"
}

func (c *GoodsTypeController) DoAdd() {
	title := models.Trim(c.GetString("title"))
	description := models.Trim(c.GetString("description"))
	status, err1 := c.GetInt("status")
	if err1 != nil {
		c.Error("参数有误", "/goodsType/add")
	}
	if title == "" {
		c.Error("标题不能为空", "/goodsType/add")
		return
	}
	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = models.GetIntTime()
	err := models.DB.Create(&goodsType).Error
	if err == nil {
		c.Success("添加角色成功", "/goodsType")
	} else {
		c.Error("添加角色失败", "/goodsType/add")
	}
}

func (c *GoodsTypeController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}

	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	c.TplName = "admin/goodsType/edit.html"
}

func (c *GoodsTypeController) DoEdit() {

	id, err1 := c.GetInt("id")
	status, err2 := c.GetInt("status")
	if err1 != nil || err2 != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}

	title := strings.Trim(c.GetString("title"), " ")
	description := strings.Trim(c.GetString("description"), " ")
	if title == "" {
		c.Error("标题不能为空", "/goodsType/add")
		return
	}
	//修改
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	err3 := models.DB.Save(&goodsType).Error
	if err3 != nil {
		c.Error("修改数据失败", "/goodsType/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/goodsType")
	}

}

func (c *GoodsTypeController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Delete(&goodsType)
	c.Success("删除角色成功", "/goodsType")

}


