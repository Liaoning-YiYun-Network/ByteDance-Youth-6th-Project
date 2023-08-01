package dao

import (
	"SkyLine/data"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"os"
)

type MySQLConfig struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *MySQLConfig) getConf() *MySQLConfig {
	//读取resources/mysql.yaml文件
	var yamlFile []byte
	var err error
	if data.OS == "windows" {
		yamlFile, err = os.ReadFile("../resources/mysql.yaml")
	} else {
		yamlFile, err = os.ReadFile("resources/mysql.yaml")
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

var SqlSession *gorm.DB

// InitMySql 初始化数据库连接
//
// 返回值：err
func InitMySql() (err error) {
	var c MySQLConfig
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
