package likeRepository

import (
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	"time"
)

type Like struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	PostId    string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (like *Like) ToEntity() *likeEntity.Like {
	return &likeEntity.Like{
		Id:        int(like.ID),
		Code:      like.Code,
		PostId:    like.PostId,
		UserId:    like.UserId,
		CreatedAt: like.CreatedAt,
		UpdatedAt: like.UpdatedAt,
	}
}

func (Like) FromEntity(like *likeEntity.Like) *Like {
	return &Like{
		ID:        uint(like.Id),
		Code:      like.Code,
		PostId:    like.PostId,
		UserId:    like.UserId,
		CreatedAt: like.CreatedAt,
		UpdatedAt: like.UpdatedAt,
	}
}
