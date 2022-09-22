package index

import (
	"github.com/astaxie/beego"
	. "github.com/hunterhug/go_image"  //前面加点，可以直接调用里面的方法
	 qrcode "github.com/skip2/go-qrcode" //使用别名
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	filename := "static/upload/a.jpg"   //原图
	savepath := "static/upload/a_1.jpg"  //编辑后存放图片

	width := 800  //编辑后宽度
	hight := 400  //高度

	//实现图片裁切
	err1 := ScaleF2F(filename, savepath, width)  //按宽度等比例缩放
	if err1 != nil {
		beego.Error(err1)
	}

	err := ThumbnailF2F(filename, savepath, width, hight)
	if err != nil {
		beego.Error(err)
	}

	err3 := qrcode.WriteFile("https://itying.com",qrcode.Medium,256,"static/upload/qr.png")  //生成二维码
	if err3 != nil {
		beego.Error(err3)
	}
	c.TplName = "index/index.html"
}
