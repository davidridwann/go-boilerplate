package logRepository

import (
	"fmt"
	logEntity "github.com/davidridwann/wlb-test.git/internal/entity/log"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	mongoConfig "github.com/davidridwann/wlb-test.git/pkg/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var logCollection *mongo.Collection = mongoConfig.GetCollection(mongoConfig.DB, "log")

func CreateLog(path string, user logEntity.User, req logEntity.Request, res logEntity.Response, ctx *gin.Context) (Log, error) {
	data := Log{
		Path:       path,
		User:       LogUser(user),
		TimeToLive: fmt.Sprintf("%.6fms", time.Since(time.Now()).Seconds()/1000),
		Request:    LogRequest(req),
		Response:   LogResponse(res),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	result, err := logCollection.InsertOne(ctx, data)
	if err != nil {
		log.Err().Error("Failed store log", err.Error())
	}

	log.Std().Infoln("Success store log", result)
	return data, nil
}
