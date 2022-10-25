package logEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Log struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Path       string             `json:"path"`
	User       User               `json:"user"`
	TimeToLive string             `json:"timeToLive"`
	Request    Request            `json:"request"`
	Response   Response           `json:"response"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type User struct {
	User string `json:"user"`
}

type Request struct {
	Request string `json:"request"`
}

type Response struct {
	Response string `json:"response"`
}
