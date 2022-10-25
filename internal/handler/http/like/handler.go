package likeHandler

import (
	"fmt"
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	likeUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/like"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type restHandler struct {
	likeUseCase likeUseCase.IUseCase
	postUseCase postUseCase.IUseCase
	userUseCase userUseCase.IUseCase
}

func New(likeUseCase likeUseCase.IUseCase, postUseCase postUseCase.IUseCase, userUseCase userUseCase.IUseCase) RestHandler {
	return &restHandler{likeUseCase, postUseCase, userUseCase}
}

// Like      	 godoc
// @Summary      Like a post
// @Description  Like a post
// @Tags         Like
// @Produce      json
// @Security	 BearerAuth
// @Param        post_id  query    string  true  "Post Code"
// @Success      200  {object} map[string]interface{}
// @Router       /post/like [post]
func (h *restHandler) Like(c *gin.Context) {
	post := c.Query("post_id")
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	claims := helpers.GetUser(c)
	convertUser, _ := json.Marshal(claims)
	convertRequest, _ := json.Marshal(post)

	if post == "" {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Post not found"})
	}

	form := likeEntity.LikeForm{
		PostId: post,
		UserId: claims.Code,
	}

	postData, err := h.postUseCase.Show(form.PostId)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Post not found"})
		return
	}

	err = h.likeUseCase.CheckLike(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{Message: "Already like a post"})
		return
	}

	err = h.likeUseCase.Like(form)
	if err == nil {
		userData, err := h.userUseCase.Get(postData.UserId)
		likeUser, err := h.userUseCase.Get(claims.Code)

		likeMessage := "New like from " + likeUser.Name

		helpers.SengNotifEmail(userData.Email, likeMessage, "New like", c)
		_, err = helpers.CreateLog("/post/like", string(convertUser), string(convertRequest), "Successfully like a post", c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Message: "Successfully like a post"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}

// Unlike      	 godoc
// @Summary      Unlike a post
// @Description  Unlike a post
// @Tags         Like
// @Produce      json
// @Security	 BearerAuth
// @Param        post_id  query    string  true  "Post Code"
// @Success      200  {object} map[string]interface{}
// @Router       /post/unlike [delete]
func (h *restHandler) Unlike(c *gin.Context) {
	post := c.Query("post_id")
	fmt.Println(post)
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	claims := helpers.GetUser(c)
	convertUser, _ := json.Marshal(claims)
	convertRequest, _ := json.Marshal(post)

	if post == "" {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Post not found"})
	}

	form := likeEntity.LikeForm{
		PostId: post,
		UserId: claims.Code,
	}

	err = h.likeUseCase.Unlike(form)
	if err == nil {
		_, err = helpers.CreateLog("/post/unlike", string(convertUser), string(convertRequest), "Successfully unlike a post", c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Message: "Successfully unlike a post"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}
