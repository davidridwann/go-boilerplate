package postRepository

import (
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	"time"
)

type Post struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	Caption   string
	Image     string
	UserId    string
	IsComment bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (post *Post) ToEntity() *postEntity.Post {
	return &postEntity.Post{
		Id:        int(post.ID),
		Code:      post.Code,
		Caption:   post.Caption,
		Image:     post.Image,
		UserId:    post.UserId,
		IsComment: post.IsComment,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (post *Post) ToEntityShow() *postEntity.PostShow {
	return &postEntity.PostShow{
		Id:        int(post.ID),
		Code:      post.Code,
		Caption:   post.Caption,
		Image:     post.Image,
		UserId:    post.UserId,
		IsComment: post.IsComment,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (Post) FromEntity(post *postEntity.Post) *Post {
	return &Post{
		ID:        uint(post.Id),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Code:      post.Code,
		Caption:   post.Caption,
		Image:     post.Image,
	}
}
