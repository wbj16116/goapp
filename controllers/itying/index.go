package itying

import (
	"goapp/models"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"strings"
	"fmt"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	// models.CacheDb.Put("username", "王宝东", time.Second*60*60)

	// username := models.CacheDb.Get("username")
	// fmt.Println("-----------------------------------")
	// fmt.Printf("%T---%v",username, username)
	// v, _ := username.([]uint8)  //类型断言 判断是不是 []uint8类型
	// fmt.Printf("%T--%v--%v",username, v, string(v))
	//获取顶部导航
	topNav := []models.Nav{}

	if hasTopNav := models.CacheDb.Get("topNav", &topNav); hasTopNav == true {
		fmt.Println("----------redis-------")
		c.Data["topNavList"] = topNav
	} else {
		models.DB.Where("status=1 AND position=1").Order("sort DESC").Find(&topNav)
		c.Data["topNavList"] = topNav

		//转为json
		// bytes, _ := json.Marshal(topNav)
		// fmt.Printf("---------------%T--------------", bytes)
		//存入redis
		models.CacheDb.Set("topNav", topNav)
	}

	//获取轮播图
	focus := []models.Focus{}
	models.DB.Where("status=1 AND focus_type=1").Order("sort DESC").Find(&focus)
	c.Data["focusList"] = focus
	//左侧分类
	goodsCate := []models.GoodsCate{}
	//models.DB.Preload("GoodsCateItem").Where("pid=0 AND status = 1").Order("sort DESC").Find(&goodsCate)  //这样只能对顶级排序
	//对顶级和子级都排序
	models.DB.Preload("GoodsCateItem", func(db *gorm.DB) *gorm.DB {
		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	}).Where("pid=0 AND status=1").Order("sort DESC").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	//获取中间导航的数据
	middleNav := []models.Nav{}
	models.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleNav)
	for i := 0; i < len(middleNav); i++ {
		//获取关联商品
		// middleNav[i].Relation  19,20,21
		middleNav[i].Relation = strings.ReplaceAll(middleNav[i].Relation, "，", ",")
		relation := strings.Split(middleNav[i].Relation, ",")
		goods := []models.Goods{}
		models.DB.Where("id in (?)", relation).Select("id,title,goods_img,price").Find(&goods)
		middleNav[i].GoodsItem = goods
	}
	//fmt.Printf("%#v", middleNav)
	c.Data["middleNavList"] = middleNav

	//获取楼层数据
	//手机
	phone := models.GetGoodsByCategory(22, "hot", 8)
	c.Data["phoneList"] = phone

	//电视
	tv := models.GetGoodsByCategory(23, "best", 8)
	c.Data["tvList"] = tv

	c.TplName = "itying/index/index.html"
}
