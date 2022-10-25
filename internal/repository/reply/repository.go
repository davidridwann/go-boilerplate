package replyRepository

import (
	"errors"
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var ErrPostNotFound = errors.New("Post not found")
var ErrUnexpected = errors.New("Unexpected Error")

type ReplyRepository interface {
	Reply(reply replyEntity.Reply) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReplyRepository {
	return &Repository{db: db}
}

func (r *Repository) Reply(reply replyEntity.Reply) error {
	code := uuid.New()

	data := Reply{
		Code:      code.String(),
		CommentId: reply.CommentId,
		Text:      reply.Text,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := r.db.Table("replies").Create(&data).Error
	if err != nil {
		log.Err().Error(err)
		return err
	}

	return nil
}
