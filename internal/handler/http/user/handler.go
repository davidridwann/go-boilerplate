package userHandler

import (
	"errors"
	"fmt"
	logEntity "github.com/davidridwann/wlb-test.git/internal/entity/log"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	logRepository "github.com/davidridwann/wlb-test.git/internal/repository/log"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
	"strings"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type restHandler struct {
	userUseCase userUseCase.IUseCase
}

func New(userUseCase userUseCase.IUseCase) RestHandler {
	return &restHandler{userUseCase}
}

func (h *restHandler) Get(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	data, err := h.userUseCase.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{
			Message: "Successfully retrieved",
			Data:    data,
		})
	} else {
		if errors.Is(err, userUseCase.ErrUnexpected) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

// Register      godoc
// @Summary      Register a user account
// @Description  Register a new user account
// @Tags         Authentication
// @Produce      json
// @Param 		 request body userEntity.User true "query params"
// @Success      200  {object} userEntity.UserAccess
// @Router       /auth/register [post]
func (h *restHandler) Register(c *gin.Context) {
	body := &userEntity.User{}
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.userUseCase.Register(*body)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result, Message: "Register Successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

// Login      godoc
// @Summary      Login a user account
// @Description  Login a user account
// @Tags         Authentication
// @Produce      json
// @Param 		 request body userEntity.AuthRequest true "query params"
// @Success      200  {object} userEntity.UserAccess
// @Router       /auth/login [post]
func (h *restHandler) Login(c *gin.Context) {
	var request userEntity.AuthRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	result, err := h.userUseCase.Login(request.Email, request.Password)
	convertRes, _ := json.Marshal(result)
	convertReq, _ := json.Marshal(request)

	var res = logEntity.Response{
		string(convertRes),
	}
	var user = logEntity.User{
		string(convertRes),
	}
	var req = logEntity.Request{
		string(convertReq),
	}

	log := logRepository.CreateLog(
		"/auth/login",
		user,
		req,
		res,
		c,
	)

	fmt.Println(log)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
	} else {
		c.JSON(http.StatusOK, SuccessResponse{
			Data:    result,
			Message: "Login Successfully",
			Success: true,
		})
	}
}

// User          godoc
// @Summary      Get user login data
// @Description  Responds with the data of user login.
// @Tags         Authentication
// @Produce      json
// @Security	 BearerAuth
// @Success      200  {object} userEntity.UserAccess
// @Router       /auth/user [get]
func (h *restHandler) User(c *gin.Context) {
	fmt.Println(c.GetHeader("Authorization"))
	token, err := jwt.ParseWithClaims(strings.Split(c.GetHeader("Authorization"), "Bearer ")[1], &helpers.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("supersecretkey"), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
	}

	if claims, ok := token.Claims.(*helpers.JWTClaim); ok && token.Valid {
		result, err := h.userUseCase.Get(claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
				Success: false,
			})
		} else {
			c.JSON(http.StatusOK, SuccessResponse{
				Data:    result,
				Message: "Login Successfully",
				Success: true,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
	}
}
