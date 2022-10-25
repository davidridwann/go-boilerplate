package postUseCase

import (
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	"github.com/gin-gonic/gin"
)

type IUseCase interface {
	Get() ([]*postEntity.PostShow, error)
	Show(code string) (*postEntity.PostShow, error)
	Create(caption string, isComment bool, user string, c *gin.Context) (*postEntity.PostShow, error)
	Update(post postEntity.PostForm, c *gin.Context) (*postEntity.PostShow, error)
	SoftDeletePost(code string) error
}
