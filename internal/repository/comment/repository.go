package commentRepository

import (
	"errors"
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var ErrPostNotFound = errors.New("Post not found")
var ErrCommentNotFound = errors.New("Comment not found")
var ErrUnexpected = errors.New("Unexpected Error")

type CommentRepository interface {
	Comment(form commentEntity.Comment) error
	Show(code string) (*commentEntity.Comment, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) CommentRepository {
	return &Repository{db: db}
}

func (r *Repository) Comment(param commentEntity.Comment) error {
	code := uuid.New()

	data := Comment{
		Code:      code.String(),
		PostId:    param.PostId,
		UserId:    param.UserId,
		Text:      param.Text,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := r.db.Table("comments").Create(&data).Error
	if err != nil {
		log.Err().Error(err)
		return err
	}

	return nil
}

func (r *Repository) Show(code string) (*commentEntity.Comment, error) {
	commentData := &Comment{}

	err := r.db.Raw(`SELECT * FROM comments
		WHERE code = ?`, code).First(&commentData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	data := commentData.ToEntity()
	return data, err
}
