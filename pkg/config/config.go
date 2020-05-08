package config

import "github.com/spf13/viper"

const (
	defaultServerHost    = "127.0.0.1"
	defaultServerPort    = 22
	defaultRedisHost     = "127.0.0.1"
	defaultRedisPort     = 6379
	defaultRedisPassword = ""
	defaultRedisDB       = 0
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.SetDefault("server", server{
		host: defaultServerHost,
		port: defaultServerPort,
	})
	viper.SetDefault("redis", redis{
		host:     defaultRedisHost,
		port:     defaultRedisPort,
		password: defaultRedisPassword,
		db:       defaultRedisDB,
	})
	viper.AddConfigPath("/etc/kobe")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
}

type server struct {
	host string
	port int
}

type redis struct {
	host     string
	port     int
	password string
	db       int
}
