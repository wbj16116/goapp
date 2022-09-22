package models

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redisClient cache.Cache
var enableRedis, _ = beego.AppConfig.Bool("enableRedis") //是否开启redis缓存 ,小写开头，只在当前文件有效
var redisTime, _ = beego.AppConfig.Int("redisTime")
func init() {
	if enableRedis {
		config := map[string]string{
			"key": beego.AppConfig.String("redisKey"),
			"conn": beego.AppConfig.String("redisConn"),
			"dbNum": beego.AppConfig.String("reidsDbNum"),
			"password":beego.AppConfig.String("redisPwd"),
		}

		bytes, _ := json.Marshal(config)  //map数组转字符串，返回byte 类型的切片
		redisClient, err = cache.NewCache("redis", string(bytes))
		if err != nil {
			fmt.Println("连接redis失败")
		} else {
			fmt.Println("连接redis数据库成功")
		}
	}
}

//定义结构体
type cacheDb struct{}

//写入数据
func (c cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value) //对象转json
		redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
	}
}

//获取数据
func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			redisValue, ok:= redisStr.([]uint8)
			if !ok {
				fmt.Println("获取redis数据失败")
			}
			json.Unmarshal([]byte(redisValue), obj)
			return true
		}
		return false
	}
	return false
}

//实例化结构体
var CacheDb = &cacheDb{}

