package admin

import (
	"goapp/models"
	"strconv"
)

type FocusController struct {
	BaseController
}

func (c *FocusController) Get() {
	focus := []models.Focus{}
	models.DB.Find(&focus)
	c.Data["focusList"] = focus
	c.TplName = "admin/focus/index.html"
}

func (c *FocusController) Add() {
	c.TplName = "admin/focus/add.html"
}

func (c *FocusController) DoAdd() {
	focusType, err1 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err1 != nil || err3 != nil {
		c.Error("非法请求", "/focus/add")
	}

	if err2 != nil {
		c.Error("排序表单里面输入的数据不合法", "/focus/add")
	}

	//执行图片上传
	focusImgSrc, _ := c.UploadImg("focus_img")

	focus := models.Focus{
		Title: title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err4 := models.DB.Create(&focus).Error
	if err4 != nil {
		c.Error("添加失败", "/focus/add")
		return
	}
	c.Success("增加轮播图成功", "/focus")
}

func (c *FocusController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/focus")
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.Data["focus"] = focus

	c.TplName = "admin/focus/edit.html"
}

func (c *FocusController) DoEdit() {
	id, err1 := c.GetInt("id")
	focusType, err2 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err3 := c.GetInt("sort")
	status, err4 := c.GetInt("status")
	if err1 != nil || err2 != nil || err4 != nil {
		c.Error("非法请求", "/focus")
	}

	if err3 != nil {
		c.Error("排序表单里面输入的数据不合法", "/focus/edit?id=" + strconv.Itoa(id))
	}

	//执行图片上传
	focusImgSrc, _ := c.UploadImg("focus_img")
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		// c.Data["err5"] = err5
		// c.ServeJSON()
		c.Error("编辑失败", "/focus/edit?id=" + strconv.Itoa(id))
		return

	}
	c.Success("编辑成功", "/focus")
}
func (c *FocusController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("ID有误", "/focus")
		return
	}
	focus := models.Focus{Id: id}
	err2 := models.DB.Delete(&focus).Error

	if err2 != nil {
		c.Error("删除失败", "/focus")
		return
	}

	c.Success("删除成功", "/focus")
}

