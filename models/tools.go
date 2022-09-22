package models

import (
	"fmt"
	"time"
	"crypto/md5"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"reflect"
	"path"
	. "github.com/hunterhug/go_image"
)



func DateToUnix(str string) int64 {

	template := "2007-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

//获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2021-01-02 15:25:04"
	t := time.Now()
	return t.Format(template)
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Trim(str string) string {
	return strings.Trim(str, "");
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

//通过反射获取商店配置中指定的字段
func GetSettingFromColumn(columnName string) string {
	setting := Setting{}
	DB.First(&setting)

	//无法用 sting类型的字段名从结构体中获取值
	//所以用反射方法来获取对应字段的值
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

//生成不同尺寸的缩略图
func ResizeImage(filename string) {
	extName := path.Ext(filename)  //文件后缀名
	sizelist := strings.Split(beego.AppConfig.String("resizeImageSize"), ",") //字符串拆分为数组
	for _, w := range sizelist {
		savepath := filename + "_" + w + "x" + w + extName //新的文件名
		width, _ := strconv.Atoi(w)
		err := ThumbnailF2F(filename, savepath , width, width)
		if err != nil {
			beego.Error(err)
		}
	}
}

//根据是否是oss图片，返回对应的路径、域名
func FormatImg(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {  //如果参数有误
		//判断目录中是否包含字符串
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		} else {
			return "/" + picName
		}
	}
}
