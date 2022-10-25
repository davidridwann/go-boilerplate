package replyRepository

import (
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	"time"
)

type Reply struct {
	ID        uint      `gorm:"primarykey"`
	Code      string    `gorm:"unique"`
	CommentId string    `json:"comment_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (reply *Reply) ToEntity() *replyEntity.Reply {
	return &replyEntity.Reply{
		Id:        int(reply.ID),
		Code:      reply.Code,
		CommentId: reply.CommentId,
		Text:      reply.Text,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
}

func (Reply) FromEntity(reply *replyEntity.Reply) *Reply {
	return &Reply{
		ID:        uint(reply.Id),
		Code:      reply.Code,
		CommentId: reply.CommentId,
		Text:      reply.Text,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
}
