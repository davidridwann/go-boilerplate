package helpers

import (
	logEntity "github.com/davidridwann/wlb-test.git/internal/entity/log"
	logRepository "github.com/davidridwann/wlb-test.git/internal/repository/log"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
)

func CreateLog(path string, account string, request string, response string, c *gin.Context) (logRepository.Log, error) {
	var res = logEntity.Response{
		Response: response,
	}
	var user = logEntity.User{
		User: string(account),
	}
	var req = logEntity.Request{
		Request: string(request),
	}

	data, err := logRepository.CreateLog(
		path,
		user,
		req,
		res,
		c,
	)

	if err != nil {
		log.Err().Fatalln("Failed store log", err.Error())
	}

	return data, nil
}
