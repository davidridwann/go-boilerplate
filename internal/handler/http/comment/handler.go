package commentHandler

import (
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	commentUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/comment"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type restHandler struct {
	commentUseCase commentUseCase.IUseCase
	postUseCase    postUseCase.IUseCase
	userUserCase   userUseCase.IUseCase
}

func New(commentUseCase commentUseCase.IUseCase, postUseCase postUseCase.IUseCase, userUseCase userUseCase.IUseCase) RestHandler {
	return &restHandler{commentUseCase: commentUseCase, postUseCase: postUseCase, userUserCase: userUseCase}
}

// Comment       godoc
// @Summary      Comment a post
// @Description  Comment a post
// @Tags         Comment
// @Produce      json
// @Security	 BearerAuth
// @Param        post_id  query    string  true  "Post Code"
// @Param        comment  query    string  true  "Comment"
// @Success      200  {object} map[string]interface{}
// @Router       /post/comment [post]
func (h *restHandler) Comment(c *gin.Context) {
	claims := helpers.GetUser(c)

	comment := commentEntity.Comment{
		PostId: c.Query("post_id"),
		Text:   c.Query("comment"),
		UserId: claims.Code,
	}
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	convertUser, _ := json.Marshal(claims)
	convertRequest, _ := json.Marshal(comment)

	if comment.PostId == "" {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Post not found"})
		return
	}

	post, err := h.postUseCase.Show(comment.PostId)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Post not found"})
		return
	}

	if post.IsComment == false {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{Message: "Comment disable for this post"})
		return
	}

	err = h.commentUseCase.Comment(comment)
	if err == nil {
		user, err := h.userUserCase.Get(post.UserId)
		if err != nil {
			log.Err().Panic(err)
		}

		helpers.SengNotifEmail(user.Email, comment.Text, "New Comment", c)

		_, err = helpers.CreateLog("/post/comment", string(convertUser), string(convertRequest), "Successfully comment a post", c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Message: "Successfully comment a post"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}
