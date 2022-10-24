package postHandler

import (
	"fmt"
	_ "github.com/davidridwann/wlb-test.git/internal/entity/post"
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type restHandler struct {
	postUseCase postUseCase.IUseCase
}

func New(postUseCase postUseCase.IUseCase) RestHandler {
	return &restHandler{postUseCase}
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
	fmt.Println(param)
	data, err := h.postUseCase.Show(param)
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
	result, err := h.postUseCase.Create(form.Caption, form.IsComment, c)
	if err == nil {
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

	c.JSON(http.StatusOK, helpers.SuccessResponse{
		Message: "Successfully deleted post",
	})
}
