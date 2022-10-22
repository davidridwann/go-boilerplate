package main

import (
	userHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/user"
	userRepository "github.com/davidridwann/wlb-test.git/internal/repository/user"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func newRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	repositoryAuth := userRepository.New(db)
	useCaseAuth := userUseCase.NewUseCase(repositoryAuth)
	handlerAuth := userHandler.New(useCaseAuth)

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", handlerAuth.Login)
		authRoutes.POST("/register", handlerAuth.Register)
	}

	return router
}
