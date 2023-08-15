package config

import (
	"SkyLine/data"
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig 使用viper读取配置文件
func InitConfig() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../resources/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件出错，请检查配置文件是否存在，运行终止！")
		//直接退出程序
		panic(err)
	}
	fmt.Println("配置文件读取成功！")
	data.DefaultSignature = viper.GetString("default.signature")
	data.DefaultAvatar = viper.GetString("default.avatar")
	data.DefaultBackgroundImage = viper.GetString("default.background_image")
}
