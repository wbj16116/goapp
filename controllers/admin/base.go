package admin

import (
	"github.com/astaxie/beego"
	"errors"
	"os"
	"path"
	"strconv"
	"goapp/models"
	"strings"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)
var adminPath string = beego.AppConfig.String("adminPath")
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, redirect string) {
	//判断路径中是否包含了域名（完整路径）
	if (strings.Contains(redirect, adminPath)){  //判断是否包含字符串
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = adminPath + redirect
	}
	c.Data["message"] = message

	c.TplName = "admin/public/success.html"
}

func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = adminPath + redirect
	c.TplName = "admin/public/error.html"
}

func (c *BaseController) ToPage(redirect string) {
	c.Redirect(adminPath + redirect, 302)
}

//上传图片
func (c *BaseController) UploadImg(picName string) (string, error) {
	ossStatus, _ := beego.AppConfig.Bool("ossStatus")
	if ossStatus == true {
		return c.OssUploadImg(picName)
	} else {
		return c.LocalUploadImg(picName)
	}
}

//本地上传
func (c *BaseController) LocalUploadImg(picName string) (string, error) {
	//1、获取上传的文件
	f, h, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {

		return "", errors.New("图片后缀名不合法")
	}
	//4、创建图片保存目录  static/upload/20200623
	day := models.GetDay()
	dir := "static/upload/" + day

	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	//5、生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(models.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	//6、保存图片
	c.SaveToFile(picName, saveDir)

	return saveDir, nil
}

//Oss上传
func (c *BaseController) OssUploadImg(picName string) (string, error) {
	//获取oss配置信息
	setting := models.Setting{}
	models.DB.First(&setting)

	//1、获取上传的文件
	f, h, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}

	//4.1、创建Oss实例
	client, err := oss.New(setting.EndPoint, setting.Appid, setting.AppSecret)
	if err != nil {
		return "", err
	}

	//4.2获取存储p空间
	bucket, err := client.Bucket(setting.BucketName)
	if err != nil {
		return "", err
	}

	//4.3创建图片保存目录
	day := models.GetDay()
	dir := "static/upload/" + day
	fileUnixName := strconv.FormatInt(models.GetUnixNano(), 10)
	saveDir := path.Join(dir, fileUnixName + extName)

	//4.4上传流文件
	err = bucket.PutObject(saveDir, f)
	if err != nil {
		return "", err
	}
	return saveDir, nil
}


func (c *BaseController) GetSetting() models.Setting {
	setting := models.Setting{}
	models.DB.First(&setting)
	return setting
}
