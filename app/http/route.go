package main

import (
	_ "github.com/davidridwann/wlb-test.git/docs"
	commentHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/comment"
	likeHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/like"
	logHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/log"
	postHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/post"
	replyHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/reply"
	userHandler "github.com/davidridwann/wlb-test.git/internal/handler/http/user"
	"github.com/davidridwann/wlb-test.git/internal/middleware"
	commentRepository "github.com/davidridwann/wlb-test.git/internal/repository/comment"
	likeRepository "github.com/davidridwann/wlb-test.git/internal/repository/like"
	postRepository "github.com/davidridwann/wlb-test.git/internal/repository/post"
	replyRepository "github.com/davidridwann/wlb-test.git/internal/repository/reply"
	userRepository "github.com/davidridwann/wlb-test.git/internal/repository/user"
	commentUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/comment"
	likeUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/like"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	replyUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/reply"
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

// @host wlb.sociolite.id
// @basePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func newRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// user
	repositoryAuth := userRepository.New(db)
	useCaseAuth := userUseCase.NewUseCase(repositoryAuth)
	handlerAuth := userHandler.New(useCaseAuth)

	// post
	repositoryPost := postRepository.New(db)
	useCasePost := postUseCase.NewUseCase(repositoryPost)
	handlerPost := postHandler.New(useCasePost, useCaseAuth)

	// like
	reposistoryLike := likeRepository.New(db)
	useCaseLike := likeUseCase.NewUseCase(reposistoryLike)
	handlerLike := likeHandler.New(useCaseLike, useCasePost, useCaseAuth)

	// comment
	repositoryComment := commentRepository.New(db)
	useCaseComment := commentUseCase.NewUseCase(repositoryComment)
	handlerComment := commentHandler.New(useCaseComment, useCasePost, useCaseAuth)

	// reply
	repositoryReply := replyRepository.New(db)
	useCaseReply := replyUseCase.NewUseCase(repositoryReply)
	handlerReply := replyHandler.New(useCaseReply, useCaseComment, useCasePost, useCaseAuth)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "Work Life & Beyond Backend TEST, Wish me luck!",
		})
	})

	api := router.Group("/api")
	{
		api.GET("/log", logHandler.Get)
		authRoutes := api.Group("auth")
		{
			authRoutes.POST("/login", handlerAuth.Login)
			authRoutes.POST("/register", handlerAuth.Register)
			authRoutes.POST("/verification-account", handlerAuth.VerifAccount)
			secured := api.Group("/auth").Use(middleware.Auth())
			{
				secured.GET("/user", handlerAuth.User)
				post := api.Group("post").Use(middleware.Auth())
				{
					// post
					post.GET("", handlerPost.Get)
					post.GET("show", handlerPost.Show)
					post.POST("create", handlerPost.Create)
					post.PUT("update", handlerPost.Update)
					post.DELETE("delete", handlerPost.Delete)

					// like
					post.POST("like", handlerLike.Like)
					post.DELETE("unlike", handlerLike.Unlike)

					// comment
					post.POST("comment", handlerComment.Comment)

					// reply
					post.POST("comment/reply", handlerReply.Reply)
				}
			}
		}
	}

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
