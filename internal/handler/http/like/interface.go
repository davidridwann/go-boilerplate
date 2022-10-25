package likeHandler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Like(c *gin.Context)
	Unlike(c *gin.Context)
}
