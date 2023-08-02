package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *RedisConfig) getConf() *RedisConfig {
	c.Url = viper.GetString("redis.url")
	c.UserName = viper.GetString("redis.userName")
	c.Password = viper.GetString("redis.password")
	c.DbName = viper.GetString("redis.dbname")
	c.Port = viper.GetString("redis.port")
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
	fmt.Printf("redis配置文件参数：%#v\n", conf)
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
