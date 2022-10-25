package postUseCase

import (
	"fmt"
	postEntity "github.com/davidridwann/wlb-test.git/internal/entity/post"
	postRepository "github.com/davidridwann/wlb-test.git/internal/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type IUseCaseImplementation struct {
	postRepository postRepository.PostRepository
}

func NewUseCase(postRepository postRepository.PostRepository) IUseCase {
	return &IUseCaseImplementation{postRepository}
}

func (uc *IUseCaseImplementation) Get() ([]*postEntity.PostShow, error) {
	data, err := uc.postRepository.Get()
	if err != nil {
		if errors.Is(err, postRepository.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Show(code string) (*postEntity.PostShow, error) {
	data, err := uc.postRepository.Show(code)
	if err != nil {
		if errors.Is(err, postRepository.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Create(caption string, isComment bool, user string, c *gin.Context) (*postEntity.PostShow, error) {
	data, err := uc.postRepository.Create(caption, isComment, user, c)
	if err != nil {
		return nil, ErrUnexpected
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Update(post postEntity.PostForm, c *gin.Context) (*postEntity.PostShow, error) {
	data, err := uc.postRepository.Update(post, c)
	if err != nil {
		return nil, ErrUnexpected
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Delete(code string) error {
	err := uc.postRepository.Delete(code)
	if err != nil {
		return err
	}

	return nil
}
