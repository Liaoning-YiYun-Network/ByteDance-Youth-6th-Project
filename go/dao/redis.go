package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type RedisConfig struct {
	Url      string
	UserName string
	Password string
	DbName   string
	Port     string
}

func (c *RedisConfig) getConf() *RedisConfig {
	c.Url = viper.GetString("redis.url")
	c.UserName = viper.GetString("redis.userName")
	c.Password = viper.GetString("redis.password")
	c.DbName = viper.GetString("redis.dbname")
	c.Port = viper.GetString("redis.port")
	return c
}

var RedisSession *redis.Client

// InitRedis 初始化Redis连接
//
// 返回值：err
func InitRedis() (err error) {
	var c RedisConfig
	conf := c.getConf()
	strBuilder := strings.Builder{}
	strBuilder.WriteString(conf.Url)
	strBuilder.WriteString(":")
	strBuilder.WriteString(conf.Port)
	db, err := strconv.Atoi(conf.DbName)
	if err != nil {
		err := fmt.Errorf("redis db name is not a number")
		return err
	}
	RedisSession = redis.NewClient(&redis.Options{
		Addr:     strBuilder.String(),
		Password: conf.Password,
		DB:       db,
	})
	pong, err := RedisSession.Ping().Result()
	if err != nil {
		err := fmt.Errorf("redis connect failed")
		return err
	}
	fmt.Println("Redis Client Connect Response:", pong)
	return nil
}

// CloseRedis 关闭Redis连接
//
// 返回值：无
func CloseRedis() {
	//关闭数据库连接
	RedisSession.Close()
}
