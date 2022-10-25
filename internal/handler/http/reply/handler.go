package replyHandler

import (
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	commentUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/comment"
	postUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/post"
	replyUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/reply"
	userUseCase "github.com/davidridwann/wlb-test.git/internal/usecase/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type restHandler struct {
	replyUseCase   replyUseCase.IUseCase
	commentUseCase commentUseCase.IUseCase
	postUseCase    postUseCase.IUseCase
	userUserCase   userUseCase.IUseCase
}

func New(replyUseCase replyUseCase.IUseCase, commentUseCase commentUseCase.IUseCase, postUseCase postUseCase.IUseCase, userUseCase userUseCase.IUseCase) RestHandler {
	return &restHandler{replyUseCase: replyUseCase, commentUseCase: commentUseCase, postUseCase: postUseCase, userUserCase: userUseCase}
}

// Reply         godoc
// @Summary      Reply a comment
// @Description  Reply a comment
// @Tags         Reply
// @Produce      json
// @Security	 BearerAuth
// @Param        comment_id  query    string  true  "Comment Code"
// @Param        reply  query    string  true  "Reply"
// @Success      200  {object} map[string]interface{}
// @Router       /post/comment/reply [post]
func (h *restHandler) Reply(c *gin.Context) {
	var commentData *commentEntity.Comment
	reply := replyEntity.Reply{
		CommentId: c.Query("comment_id"),
		Text:      c.Query("reply"),
	}
	err := c.ShouldBind(&reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	claims := helpers.GetUser(c)
	convertUser, _ := json.Marshal(claims)
	convertRequest, _ := json.Marshal(reply)

	if reply.CommentId == "" {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Comment not found"})
		return
	}

	commentData, err = h.commentUseCase.Show(reply.CommentId)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{Message: "Comment not found"})
		return
	}

	err = h.replyUseCase.Reply(reply)
	if err == nil {
		userData, err := h.userUserCase.Get(commentData.UserId)

		helpers.SengNotifEmail(userData.Email, reply.Text, "New Reply", c)
		_, err = helpers.CreateLog("/post/comment/reply", string(convertUser), string(convertRequest), "Successfully reply a comment", c)
		if err != nil {
			log.Err().Panic(err)
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{Message: "Successfully reply a comment"})
	} else {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: err.Error()})
	}
}
