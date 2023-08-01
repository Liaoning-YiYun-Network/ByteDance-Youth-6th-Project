package dao

import (
	"SkyLine/data"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"os"
)

type RedisConfig struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *RedisConfig) getConf() *RedisConfig {
	//读取resources/redis.yaml文件
	var yamlFile []byte
	var err error
	if data.OS == "windows" {
		yamlFile, err = os.ReadFile("../resources/redis.yaml")
	} else {
		yamlFile, err = os.ReadFile("resources/redis.yaml")
	}
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

var RedisSession *gorm.DB

// InitRedis 初始化Redis连接
//
// 返回值：err
func InitRedis() (err error) {
	var c RedisConfig
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	RedisSession, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return RedisSession.DB().Ping()
}

// CloseRedis 关闭Redis连接
//
// 返回值：无
func CloseRedis() {
	//关闭数据库连接
	RedisSession.Close()
}
