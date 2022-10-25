package commentEntity

import (
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	"time"
)

type Comment struct {
	Id        int       `json:"id" swaggerignore:"true"`
	Code      string    `json:"code" swaggerignore:"true"`
	PostId    string    `json:"post_id" swaggerignore:"true"`
	UserId    string    `json:"user_id"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
}

type CommentWithReply struct {
	Id        int                 `json:"id"`
	Code      string              `json:"code"`
	PostId    string              `json:"post_id"`
	UserId    string              `json:"user_id"`
	Text      string              `json:"text"`
	Replies   []replyEntity.Reply `json:"replies"`
	CreatedAt time.Time           `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time           `json:"updated_at" swaggerignore:"true"`
}
