package postEntity

import (
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	"mime/multipart"
	"time"
)

type Post struct {
	Id        int       `json:"id" swaggerignore:"true"`
	Code      string    `json:"code" swaggerignore:"true"`
	Caption   string    `json:"caption" binding:"required"`
	Image     string    `json:"image" swaggerignore:"true"`
	UserId    string    `json:"user_id" swaggerignore:"true"`
	IsComment bool      `json:"is_comment"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
	DeletedAt time.Time `json:"deleted_at" swaggerignore:"true"`
}

type PostShow struct {
	Id        int                              `json:"id"`
	Code      string                           `json:"code"`
	Caption   string                           `json:"caption"`
	Image     string                           `json:"image"`
	UserId    string                           `json:"user_id"`
	IsComment bool                             `json:"is_comment"`
	Likes     []likeEntity.Like                `json:"likes"`
	Comments  []commentEntity.CommentWithReply `json:"comments"`
	CreatedAt time.Time                        `json:"created_at"`
	UpdatedAt time.Time                        `json:"updated_at"`
}

type PostForm struct {
	Code      string                  `form:"code"`
	Caption   string                  `form:"caption" binding:"required"`
	Image     []*multipart.FileHeader `form:"image" binding:"required"`
	IsComment bool                    `form:"is_comment"`
}
