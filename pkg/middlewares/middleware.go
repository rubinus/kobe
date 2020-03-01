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
	s := db.Session.Clone()
	defer s.Close()
	d := s.DB(db.Mongo.Database)
	collection := d.C("worker")
	_ = collection.Remove(bson.M{})
	go pickAndCloseWorker()
}

func pickAndCloseWorker() {
	go func() {
		for {
			s := db.Session.Clone()
			defer s.Close()
			d := s.DB(db.Mongo.Database)
			c := d.C("worker")
			time.Sleep(30 * time.Second)
			log.Debug("start health check...")
			var workers = []models.Worker{}
			err := c.Find(bson.M{"state": models.StateWorkerOnline}).All(&workers)
			if err != nil {
				log.Error(err.Error())
			}
			log.Debugf("online worker num: %d", len(workers))
			for _, worker := range workers {
				subM := time.Now().Sub(worker.LastHealthCheckTime)
				if subM.Seconds() > 30 {
					worker.State = models.StateWorkerOffline
					if err := c.Update(bson.M{"uid": worker.Uid}, &worker); err != nil {
						log.Warn(err.Error())
					}

				}
			}
		}
	}()
}
