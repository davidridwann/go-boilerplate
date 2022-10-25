package likeRepository

import (
	"errors"
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var ErrPostNotFound = errors.New("Post not found")
var ErrUnexpected = errors.New("Unexpected Error")

type LikeRepository interface {
	Like(form likeEntity.LikeForm) error
	Unlike(form likeEntity.LikeForm) error
	CheckLike(form likeEntity.LikeForm) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) LikeRepository {
	return &Repository{db: db}
}

func (r *Repository) Like(param likeEntity.LikeForm) error {
	code := uuid.New()

	data := Like{
		Code:      code.String(),
		PostId:    param.PostId,
		UserId:    param.UserId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := r.db.Table("likes").Create(&data).Error
	if err != nil {
		log.Err().Error(err)
		return err
	}

	return nil
}

func (r *Repository) Unlike(param likeEntity.LikeForm) error {
	err := r.db.Where("post_id = ? AND user_id = ?", param.PostId, param.UserId).Delete(&Like{}).Error

	if err != nil {
		log.Err().Error(err)
		return err
	}

	return nil
}

func (r *Repository) CheckLike(param likeEntity.LikeForm) error {
	var like int64

	err := r.db.Table("likes").Where("post_id = ? AND user_id = ?", param.PostId, param.UserId).Count(&like).Error
	if err != nil {
		log.Err().Error(err)
		return err
	}

	if like != 0 {
		return errors.New("")
	}

	return nil
}
