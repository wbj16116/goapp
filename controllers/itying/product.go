package itying

import (
	"math"
	"strconv"
	"goapp/models"
	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get() {
	c.TplName = "itying/product/list.html"
}

//根据分类，获取分类下的商品
func (c *ProductController) CategoryList() {
	//获取动态路由的自定义参数  category_8.html
	id := c.Ctx.Input.Param(":id")
	cateId, _ := strconv.Atoi(id) //string 转 int
	curretGoodsCate := models.GoodsCate{}  //当前分类
	subGoodsCate := []models.GoodsCate{}   //子类
	models.DB.Where("id=?", id).Find(&curretGoodsCate)

	//分页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 5

	var tempSlice []int
	if curretGoodsCate.Pid == 0 { //如果是顶级分类
		//查询二级分类
		models.DB.Where("pid=?", curretGoodsCate.Id).Find(&subGoodsCate)
		for i := 0; i < len(subGoodsCate); i++ {
			tempSlice = append(tempSlice, subGoodsCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	goods := []models.Goods{}
	models.DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Offset((page-1) * pageSize).Limit(pageSize).Order("sort desc").Find(&goods)
	//查询goods表里的数量
	var count int
	models.DB.Where(where,tempSlice).Table("goods").Count(&count)

	c.Data["goodsList"] = goods
	c.Data["subGoodsCate"] = subGoodsCate
	c.Data["curretGoodsCate"] = curretGoodsCate
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	tpl := curretGoodsCate.Template
	if tpl == "" {
		tpl = "itying/product/list.html"
	}
	c.TplName = tpl
}
