package userHandler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	User(c *gin.Context)
}
