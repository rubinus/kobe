package worker

import (
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/db"
	"kobe/pkg/logger"
	"kobe/pkg/models"
)

var log = logger.Logger

func Run() {
	db.Connect()
	s := db.Session.Clone()
	defer s.Close()
	d := s.DB(db.Mongo.Database)
	worker := models.Worker{
		Db:    d,
		State: models.StateWorkerOnline,
		Id:    uuid.NewV4().String(),
	}
	worker.Listen()
}
