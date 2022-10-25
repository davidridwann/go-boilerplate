package postRepository

import (
	"errors"
	"fmt"
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

var ErrPostNotFound = errors.New("Post not found")
var ErrUnexpected = errors.New("Unexpected Error")
var ErrDisableComment = errors.New("Comment disable for this post")

type PostRepository interface {
	Get() ([]*postEntity.PostShow, error)
	Show(code string) (*postEntity.PostShow, error)
	Create(caption string, isComment bool, user string, c *gin.Context) (*postEntity.PostShow, error)
	Update(post postEntity.PostForm, c *gin.Context) (*postEntity.PostShow, error)
	SoftDeletePost(code string) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) PostRepository {
	return &Repository{db: db}
}

func (r *Repository) Get() ([]*postEntity.PostShow, error) {
	var postData []Post
	var likeData []likeEntity.Like
	var commentData []commentEntity.Comment
	var commentWithReply []commentEntity.CommentWithReply
	var replyData []replyEntity.Reply

	err := r.db.Raw(`SELECT * FROM posts WHERE deleted_at IS NULL`).Find(&postData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	var data []*postEntity.PostShow
	for _, v := range postData {
		err = r.db.Raw(`SELECT * FROM likes WHERE post_id = ?`, v.Code).Find(&likeData).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Err().Error(err)
				return nil, err
			}
		}

		err = r.db.Raw(`SELECT * FROM comments WHERE post_id = ?`, v.Code).Find(&commentData).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Err().Error(err)
				return nil, err
			}
		}

		for _, i := range commentData {
			fmt.Printf(i.Code)
			err = r.db.Raw(`SELECT * FROM replies WHERE comment_id = ?`, i.Code).Find(&replyData).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Err().Error(err)
					return nil, err
				}
			}

			commentWithReply = append(commentWithReply, commentEntity.CommentWithReply{
				Id:        i.Id,
				Code:      i.Code,
				PostId:    i.PostId,
				Text:      i.Text,
				Replies:   replyData,
				CreatedAt: i.CreatedAt,
				UpdatedAt: i.UpdatedAt,
			})
		}

		data = append(data, &postEntity.PostShow{
			Id:        int(v.ID),
			Code:      v.Code,
			Caption:   v.Caption,
			IsComment: v.IsComment,
			Image:     v.Image,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Likes:     likeData,
			Comments:  commentWithReply,
		})
	}
	return data, err
}

func (r *Repository) Show(code string) (*postEntity.PostShow, error) {
	postData := &Post{}

	err := r.db.Raw(`SELECT * FROM posts
		WHERE code = ?`, code).First(&postData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	data := postData.ToEntityShow()
	return data, err
}

func (r *Repository) Create(caption string, isComment bool, user string, c *gin.Context) (*postEntity.PostShow, error) {
	var imageName string
	code := uuid.New()

	file, err := c.FormFile("image")
	if err != nil {
		log.Err().Error(err)
		return nil, err
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	pwd, err := os.Getwd()
	if err != nil {
		log.Err().Error(err)
		os.Exit(1)
	}
	exPath := filepath.Dir(pwd)

	imageName = filepath.Dir(exPath) + "/image/" + newFileName
	if err = c.SaveUploadedFile(file, imageName); err != nil {
		log.Err().Error(err)
		return nil, err
	}

	data := Post{
		Code:      code.String(),
		Caption:   caption,
		Image:     imageName,
		UserId:    user,
		IsComment: isComment,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = r.db.Table("posts").Create(&data).Error
	if err != nil {
		log.Err().Error(err)
		return nil, err
	}

	post, _ := r.Show(code.String())
	return post, nil
}

func (r *Repository) Update(param postEntity.PostForm, c *gin.Context) (*postEntity.PostShow, error) {
	imageName := ""
	file, err := c.FormFile("image")
	if err == nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		pwd, err := os.Getwd()
		if err != nil {
			log.Err().Error(err)
			os.Exit(1)
		}
		exPath := filepath.Dir(pwd)

		imageName = filepath.Dir(exPath) + "/image/" + newFileName
		if err = c.SaveUploadedFile(file, imageName); err != nil {
			log.Err().Error(err)
			return nil, err
		}
	}

	if imageName == "" {
		err = r.db.Table("posts").Where("code = ?", param.Code).Updates(map[string]interface{}{
			"caption":    param.Caption,
			"is_comment": param.IsComment,
			"updated_at": time.Now(),
		}).Error
	} else {
		err = r.db.Table("posts").Where("code = ?", param.Code).Updates(map[string]interface{}{
			"caption":    param.Caption,
			"image":      imageName,
			"is_comment": param.IsComment,
			"updated_at": time.Now(),
		}).Error
	}

	if err != nil {
		log.Err().Error(err)
		return nil, err
	}

	post, _ := r.Show(param.Code)
	return post, nil
}

func (r *Repository) SoftDeletePost(code string) error {
	err := r.db.Table("posts").Where("code = ?", code).Update("deleted_at", time.Now()).Error
	if err != nil {
		log.Err().Error(err)
		return err
	}

	return nil
}
