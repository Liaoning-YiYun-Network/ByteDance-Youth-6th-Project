package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 使用viper读取配置文件
func InitConfig() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../resources/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件出错")
	}
	fmt.Println("配置文件读取成功")
}
