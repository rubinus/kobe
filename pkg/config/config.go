package config

import "github.com/spf13/viper"

const (
	defaultServerHost = "127.0.0.1"
	defaultServerPort = 8080
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.SetDefault("server", server{
		host: defaultServerHost,
		port: defaultServerPort,
	})
	viper.AddConfigPath("/etc/kobe")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
}

type server struct {
	host string
	port int
}
