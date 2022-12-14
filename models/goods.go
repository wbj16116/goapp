package models

import (
	_ "github.com/jinzhu/gorm"
)

type Goods struct {
	Id            int
	Title         string
	SubTitle      string
	GoodsSn       string
	CateId        int
	ClickCount    int
	GoodsNumber   int
	Price         float64
	MarketPrice   float64
	RelationGoods string
	GoodsAttr     string
	GoodsVersion  string
	GoodsImg      string
	GoodsGift     string
	GoodsFitting  string
	GoodsColor    string
	GoodsKeywords string
	GoodsDesc     string
	GoodsContent  string
	IsDelete      int
	IsHot         int
	IsBest        int
	IsNew         int
	GoodsTypeId   int
	Sort          int
	Status        int
	AddTime       int
}

func (Goods) TableName() string {
	return "goods"
}

func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []Goods {
	goods := []Goods{}
	goodsCate := []GoodsCate{}
	DB.Where("pid=?", cateId).Find(&goodsCate)
	var tempSlice []int
	if len(goodsCate) > 0 {
		for _, v := range goodsCate {
			tempSlice = append(tempSlice, v.Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		break
	}
	DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort DESC").Find(&goods)
	return goods
}
