package commentUseCase

import (
	"errors"
	"fmt"
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
	commentRepository "github.com/davidridwann/wlb-test.git/internal/repository/comment"
	"github.com/davidridwann/wlb-test.git/pkg/log"
)

type IUseCaseImplementation struct {
	commentRepository commentRepository.CommentRepository
}

func NewUseCase(commentRepository commentRepository.CommentRepository) IUseCase {
	return &IUseCaseImplementation{commentRepository}
}

func (uc *IUseCaseImplementation) Show(code string) (*commentEntity.Comment, error) {
	data, err := uc.commentRepository.Show(code)
	if err != nil {
		if errors.Is(err, ErrCommentNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Comment(comment commentEntity.Comment) error {
	if comment.PostId == "" {
		log.Err().Error(ErrPostNotFound)
		return nil
	}

	err := uc.commentRepository.Comment(comment)
	if err != nil {
		return err
	}

	return nil
}
