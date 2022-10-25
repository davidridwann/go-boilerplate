package postHandler

import (
	_ "github.com/davidridwann/wlb-test.git/internal/entity/post"
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"net/http"
)

type restHandler struct {
	postUseCase postUseCase.IUseCase
	useCaseUser userUseCase.IUseCase
}

func New(postUseCase postUseCase.IUseCase, useCaseUser userUseCase.IUseCase) RestHandler {
	return &restHandler{postUseCase, useCaseUser}
}

// Get      	 godoc
// @Summary      Get all post
// @Description  Get all post
// @Tags         Post
// @Produce      json
// @Security	 BearerAuth
// @Success      200  {object} postEntity.PostShow
// @Router       /post [get]
func (h *restHandler) Get(c *gin.Context) {
	data, err := h.postUseCase.Get()
	if err == nil {
		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Message: "Successfully retrieved",
			Data:    data,
		})
	} else {
		if errors.Is(err, postUseCase.ErrUnexpected) {
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
		}
	}
}

// Show      	 godoc
// @Summary      Show a specific post
// @Description  Show a specific post
// @Tags         Post
// @Produce      json
// @Security	 BearerAuth
// @Param        code  query    string  true  "Post Code"
// @Success      200  {object} postEntity.PostShow
// @Router       /post/show [get]
func (h *restHandler) Show(c *gin.Context) {
	param := c.Query("code")
	data, err := h.postUseCase.Show(param)
	if err == nil {
		claims := helpers.GetUser(c)
		convertUser, _ := json.Marshal(claims)
		convertResponse, _ := json.Marshal(data)
		_, err = helpers.CreateLog("/post/show", string(convertUser), param, string(convertResponse), c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Message: "Successfully retrieved",
			Data:    data,
		})
	} else {
		if errors.Is(err, postUseCase.ErrUnexpected) {
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
		}
	}
}

// Create      	 godoc
// @Summary      Create a new post
// @Description  Create a new post
// @Tags         Post
// @Produce      mpfd
// @Security	 BearerAuth
// @Param        caption  formData    string  true  "Post Caption"
// @Param        is_comment  formData    bool  true  "Post Enable/Disable"
// @Param        image  formData    file  true  "Post Image"
// @Success      200  {object} postEntity.Post
// @Router       /post/create [post]
func (h *restHandler) Create(c *gin.Context) {
	var form postEntity.PostForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	claims := helpers.GetUser(c)

	result, err := h.postUseCase.Create(form.Caption, form.IsComment, claims.Code, c)
	if err == nil {
		convertUser, _ := json.Marshal(claims)
		convertRequest, _ := json.Marshal(form)
		convertResponse, _ := json.Marshal(result)
		_, err = helpers.CreateLog("/post/create", string(convertUser), string(convertRequest), string(convertResponse), c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Data: result, Message: "Successfully create new post"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}

// Update      	 godoc
// @Summary      Update an existing post
// @Description  Update an existing post
// @Tags         Post
// @Produce      mpfd
// @Security	 BearerAuth
// @Param        code  formData    string  true  "Existing Post Code"
// @Param        caption  formData    string  true  "Post Caption"
// @Param        is_comment  formData    bool  true  "Post Enable/Disable"
// @Param        image  formData    file  true  "Post Image"
// @Success      200  {object} postEntity.Post
// @Router       /post/update [put]
func (h *restHandler) Update(c *gin.Context) {
	var form postEntity.PostForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	result, err := h.postUseCase.Update(form, c)
	if err == nil {
		claims := helpers.GetUser(c)
		convertUser, _ := json.Marshal(claims)
		convertRequest, _ := json.Marshal(form)
		convertResponse, _ := json.Marshal(result)
		_, err = helpers.CreateLog("/post/update", string(convertUser), string(convertRequest), string(convertResponse), c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Data: result, Message: "Successfully create new post"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}

// SoftDeletePost      	 godoc
// @Summary      Soft Delete an existing post
// @Description  Soft Delete an existing post
// @Tags         Post
// @Produce      mpfd
// @Security	 BearerAuth
// @Param        code  formData    string  true  "Existing Post Code"
// @Success      200  {object} map[string]interface{}
// @Router       /post/delete [delete]
func (h *restHandler) SoftDeletePost(c *gin.Context) {
	param := c.Param("code")

	err := h.postUseCase.SoftDeletePost(param)
	if err != nil {
		if errors.Is(err, postUseCase.ErrUnexpected) {
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
		}
	}

	claims := helpers.GetUser(c)
	convertUser, _ := json.Marshal(claims)
	_, err = helpers.CreateLog("/post/delete", string(convertUser), param, "Successfully deleted post", c)
	if err != nil {
		log.Err().Panic(err)
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse{
		Message: "Successfully deleted post",
	})
}
