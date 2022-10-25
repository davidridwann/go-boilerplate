package likeUseCase

import (
	likeEntity "github.com/davidridwann/wlb-test.git/internal/entity/like"
)

type IUseCase interface {
	Like(form likeEntity.LikeForm) error
	Unlike(form likeEntity.LikeForm) error
	CheckLike(form likeEntity.LikeForm) error
}
