package commentRepository

import (
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	PostId    string
	UserId    string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (comment *Comment) ToEntity() *commentEntity.Comment {
	return &commentEntity.Comment{
		Id:        int(comment.ID),
		Code:      comment.Code,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func (Comment) FromEntity(comment *commentEntity.Comment) *Comment {
	return &Comment{
		ID:        uint(comment.Id),
		Code:      comment.Code,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
