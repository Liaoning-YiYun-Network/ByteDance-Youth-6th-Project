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

// SetRedis 向Redis中存入数据
//
// 参数：key string, value string
//
// 返回值：err
func SetRedis(key string, value string) (err error) {
	err = RedisSession.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRedis 从Redis中获取数据
//
// 参数：key string
//
// 返回值：value string, err
func GetRedis(key string) (value string, err error) {
	value, err = RedisSession.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// DelRedis 从Redis中删除数据
//
// 参数：key string
//
// 返回值：err
func DelRedis(key string) (err error) {
	err = RedisSession.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
