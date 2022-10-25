package commentUseCase

import (
	commentEntity "github.com/davidridwann/wlb-test.git/internal/entity/comment"
)

type IUseCase interface {
	Show(code string) (*commentEntity.Comment, error)
	Comment(comment commentEntity.Comment) error
}
