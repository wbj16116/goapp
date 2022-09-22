package admin

import (
	"goapp/models"
	"math"
	"strconv"
)

type NavController struct {
	BaseController
}

func (c *NavController) Get() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page =1
	}
	//每页显示数量
	pageSize := 3

	//查询数据
	nav := []models.Nav{}
	models.DB.Offset((page-1) * pageSize).Limit(pageSize).Find(&nav)

	//获取总数量
	var count int
	models.DB.Table("nav").Count(&count)

	c.Data["navList"] = nav
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName ="admin/nav/index.html"
}

func (c *NavController) Add() {
	c.TplName = "admin/nav/add.html"
}

func (c *NavController) DoAdd() {
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")

	//填充模型
	nav := models.Nav{
		Title: title,
		Link: link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}

	err := models.DB.Create(&nav).Error
	if err == nil {
		c.Success("添加成功", "/nav")
	} else {
		c.Error("添加失败", "/nav/add")
	}

}

func (c *NavController) Edit() {
	id, _ := c.GetInt("id")
	nav := models.Nav{Id:id}
	models.DB.Find(&nav)
	c.Data["nav"] = nav
	c.Data["prePage"] = c.Ctx.Request.Referer()
	c.TplName = "admin/nav/edit.html"
}

func (c *NavController) DoEdit() {
	id, _ := c.GetInt("id")
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("is_opennew")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	prevPage := c.GetString("prevPage")

	nav := models.Nav{Id:id}
	models.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status

	err := models.DB.Save(&nav).Error
	if err != nil {
		c.Error("修改失败", "/nav/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", prevPage)
	}

}

func (c *NavController) Delete() {
	id, _ := c.GetInt("id")
	nav := models.Nav{Id:id}
	models.DB.Delete(&nav)
	c.Success("删除数据成功", c.Ctx.Request.Referer())
}
