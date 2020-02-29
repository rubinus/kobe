package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"kobe/pkg/db"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"time"
)

var log = logger.Logger

func Connect(c *gin.Context) {
	s := db.Session.Clone()
	defer s.Close()
	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

func WorkerManager(c *gin.Context) {
	c.Next()
	go func() {
		for {
			time.Sleep(30 * time.Second)
			log.Debug("start health check...")
			s := db.Session.Clone()
			d := s.DB(db.Mongo.Database)
			c := d.C("worker")
			var workers []models.Worker
			err := c.Find(bson.M{"state": models.StateWorkerOnline}).All(&workers)
			if err != nil {
				log.Error(err.Error())
			}
			for _, worker := range workers {
				subM := time.Now().Sub(worker.LastHealthCheckTime)
				if subM.Minutes() > 5 {
					worker.State = models.StateWorkerOffline
					if err := c.Update(bson.M{"_id": worker.Id}, &worker); err != nil {
						log.Warn(err.Error())
					}

				}
			}
			s.Close()
			log.Debug("end health check")
		}
	}()
}
