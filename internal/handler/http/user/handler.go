package userHandler

import (
	"errors"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (h *restHandler) Register(c *gin.Context) {
	body := &userEntity.User{}
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.userUseCase.Register(*body)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result, Message: "Registrasi Berhasil"})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

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

func (h *restHandler) User(c *gin.Context) {
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
