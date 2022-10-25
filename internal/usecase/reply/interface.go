package replyUseCase

import replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"

type IUseCase interface {
	Reply(reply replyEntity.Reply) error
}
