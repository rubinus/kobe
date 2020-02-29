package models

import "time"

type Playbook struct {
	Id          string    `json:"-" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}
