package userHandler

import (
	"errors"
	logEntity "github.com/davidridwann/wlb-test.git/internal/entity/log"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	logRepository "github.com/davidridwann/wlb-test.git/internal/repository/log"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goccy/go-json"
	"net/http"
)

type restHandler struct {
	userUseCase userUseCase.IUseCase
}

func New(userUseCase userUseCase.IUseCase) RestHandler {
	return &restHandler{userUseCase}
}

func (h *restHandler) Get(c *gin.Context) {
	param := c.Param("code")

	data, err := h.userUseCase.Get(param)
	if err == nil {
		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Message: "Successfully retrieved",
			Data:    data,
		})
	} else {
		if errors.Is(err, userUseCase.ErrUnexpected) {
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
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
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.userUseCase.Register(*body)
	if err == nil {
		helpers.SendActivationMail(result.Email, result.Code, c)

		c.JSON(http.StatusOK, helpers.SuccessResponse{Data: result, Message: "Register Successfully, please check your email"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}

// Login         godoc
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
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	result, err := h.userUseCase.Login(request.Email, request.Password)
	if result.IsActive == false {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Message: "Account is not active, please verif your account",
			Success: false,
		})
		return
	}
	// Store log
	convertRes, _ := json.Marshal(result)
	convertReq, _ := json.Marshal(request)

	var res = logEntity.Response{
		Response: string(convertRes),
	}
	var user = logEntity.User{
		User: string(convertRes),
	}
	var req = logEntity.Request{
		Request: string(convertReq),
	}

	_, _ = logRepository.CreateLog(
		"/auth/login",
		user,
		req,
		res,
		c,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
	} else {
		c.JSON(http.StatusOK, helpers.SuccessResponse{
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
	claims := helpers.GetUser(c)
	result, err := h.userUseCase.Get(claims.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
			Success: false,
		})
	} else {
		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Data:    result,
			Message: "Login Successfully",
			Success: true,
		})
	}
}

// VerifAccount User          godoc
// @Summary      Verification user account
// @Description  Verification user account
// @Tags         Authentication
// @Produce      json
// @Param        verif  query    string  true  "Verification token"
// @Success      200  {object} map[string]interface{}
// @Router       /auth/verification-account [post]
func (h *restHandler) VerifAccount(c *gin.Context) {
	token := c.Query("verif")
	err := c.ShouldBind(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.userUseCase.VerifAccount(token)
	if err == nil {
		c.JSON(http.StatusOK, helpers.SuccessResponse{Message: "Successfully verif a account"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}
