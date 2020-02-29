package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
)

var (
	Session *mgo.Session
	Mongo   *mgo.DialInfo
)

const defaultMongoUrl = "mongodb://localhost:27017/kobe"

func Connect() {
	host := viper.GetString("server.db.host")
	port := viper.GetInt("server.db.port")
	name := viper.GetString("server.db.name")
	url := fmt.Sprintf("mongodb://%s:%d/%s", host, port, name)
	log.Println(url)
	mongo, err := mgo.ParseURL(defaultMongoUrl)
	if err != nil {
		fmt.Println(err)
	}
	s, err := mgo.Dial(defaultMongoUrl)
	if err != nil {
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected", url)
	Session = s
	Mongo = mongo
}
