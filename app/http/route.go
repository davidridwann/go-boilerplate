package main

import (
	_ "github.com/davidridwann/wlb-test.git/docs"
	userHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/user"
	"github.com/davidridwann/wlb-test.git/internal/middleware"
	userRepository "github.com/davidridwann/wlb-test.git/internal/repository/user"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

// @title WorkLife&Beyond BackEnd TEST API Docs
// @version 1.0
// @description BackEnd TEST API Documentations

// @host localhost:8000
// @basePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func newRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	repositoryAuth := userRepository.New(db)
	useCaseAuth := userUseCase.NewUseCase(repositoryAuth)
	handlerAuth := userHandler.New(useCaseAuth)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "Work Life & Beyond Backend TEST, Wish me luck!",
		})
	})

	api := router.Group("/api")
	{
		authRoutes := api.Group("auth")
		{
			authRoutes.POST("/login", handlerAuth.Login)
			authRoutes.POST("/register", handlerAuth.Register)
			secured := api.Group("/auth").Use(middleware.Auth())
			{
				secured.GET("/user", handlerAuth.User)
			}
		}
	}

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
