package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *MySQLConfig) getConf() *MySQLConfig {
	c.Url = viper.GetString("db.url")
	c.UserName = viper.GetString("db.userName")
	c.Password = viper.GetString("db.password")
	c.DbName = viper.GetString("db.dbname")
	c.Port = viper.GetString("db.port")
	return c
}

var SqlSession *gorm.DB

// InitMySql 初始化数据库连接
//
// 返回值：err
func InitMySql() (err error) {
	var c MySQLConfig
	//获取yaml配置参数
	conf := c.getConf()
	fmt.Printf("数据库配置文件参数：%#v\n", conf)
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	SqlSession, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return SqlSession.DB().Ping()
}

// CloseMySql 关闭数据库连接
//
// 返回值：无
func CloseMySql() {
	//关闭数据库连接
	SqlSession.Close()
}
