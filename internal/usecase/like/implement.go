package likeUseCase

import (
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
	likeRepository "github.com/davidridwann/wlb-test.git/internal/repository/like"
	"github.com/davidridwann/wlb-test.git/pkg/log"
)

type IUseCaseImplementation struct {
	likeRepository likeRepository.LikeRepository
}

func NewUseCase(likeRepository likeRepository.LikeRepository) IUseCase {
	return &IUseCaseImplementation{likeRepository}
}

func (uc *IUseCaseImplementation) Like(form likeEntity.LikeForm) error {
	if form.PostId == "" {
		log.Err().Error(ErrPostNotFound)
		return nil
	}

	err := uc.likeRepository.Like(form)
	if err != nil {
		return err
	}

	return nil
}

func (uc *IUseCaseImplementation) Unlike(form likeEntity.LikeForm) error {
	if form.PostId == "" {
		log.Err().Error(ErrPostNotFound)
		return ErrPostNotFound
	}

	err := uc.likeRepository.Unlike(form)
	if err != nil {
		return err
	}

	return nil
}

func (uc *IUseCaseImplementation) CheckLike(form likeEntity.LikeForm) error {
	if form.PostId == "" {
		log.Err().Error(ErrPostNotFound)
		return ErrPostNotFound
	}

	err := uc.likeRepository.CheckLike(form)
	if err != nil {
		return err
	}

	return nil
}
