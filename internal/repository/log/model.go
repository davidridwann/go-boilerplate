package logRepository

import (
	logEntity "github.com/davidridwann/wlb-test.git/internal/entity/log"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogUser logEntity.User
type LogRequest logEntity.Request
type LogResponse logEntity.Response

type Log struct {
	Id         primitive.ObjectID `gorm:"primarykey"`
	Path       string
	User       LogUser
	TimeToLive string
	Request    LogRequest
	Response   LogResponse
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (log *Log) ToEntity() *logEntity.Log {
	return &logEntity.Log{
		Id:         log.Id,
		Path:       log.Path,
		User:       logEntity.User(log.User),
		TimeToLive: log.TimeToLive,
		Request:    logEntity.Request(log.Request),
		Response:   logEntity.Response(log.Response),
		CreatedAt:  log.CreatedAt,
		UpdatedAt:  log.UpdatedAt,
	}
}

func (d *LogUser) UnmarshalJSONUser(bs []byte) error {
	var s logEntity.User
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	*d = LogUser(s)
	return nil
}

func (d *LogRequest) UnmarshalJSONRequest(bs []byte) error {
	var s logEntity.Request
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	*d = LogRequest(s)
	return nil
}

func (d *LogResponse) UnmarshalJSONResponse(bs []byte) error {
	var s logEntity.Response
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	*d = LogResponse(s)
	return nil
}
