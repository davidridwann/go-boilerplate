package postHandler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Get(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	SoftDeletePost(c *gin.Context)
}
