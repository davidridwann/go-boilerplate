package likeEntity

import "time"

type Like struct {
	Id        int       `json:"id" swaggerignore:"true"`
	Code      string    `json:"code" swaggerignore:"true"`
	PostId    string    `json:"post_id" swaggerignore:"true"`
	UserId    string    `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
}

type LikeForm struct {
	Code   string `json:"code"`
	PostId string `json:"post_id" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}
