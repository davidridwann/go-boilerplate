package replyEntity

import "time"

type Reply struct {
	Id        int       `json:"id" swaggerignore:"true"`
	Code      string    `json:"code" swaggerignore:"true"`
	CommentId string    `json:"comment_id" swaggerignore:"true"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
}
