package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/etc/kobe")
	_ = viper.ReadInConfig()
}
