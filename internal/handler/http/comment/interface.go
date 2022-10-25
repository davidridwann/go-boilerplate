package commentHandler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Comment(c *gin.Context)
}
