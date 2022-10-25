package replyHandler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Reply(c *gin.Context)
}
