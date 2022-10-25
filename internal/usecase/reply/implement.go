package replyUseCase

import (
	replyEntity "github.com/davidridwann/wlb-test.git/internal/entity/reply"
	replyRepository "github.com/davidridwann/wlb-test.git/internal/repository/reply"
	"github.com/davidridwann/wlb-test.git/pkg/log"
)

type IUseCaseImplement struct {
	replyRepository replyRepository.ReplyRepository
}

func NewUseCase(replyRepository replyRepository.ReplyRepository) IUseCase {
	return &IUseCaseImplement{replyRepository: replyRepository}
}

func (uc *IUseCaseImplement) Reply(reply replyEntity.Reply) error {
	if reply.CommentId == "" {
		log.Err().Error(ErrCommentNotFound)
		return nil
	}

	err := uc.replyRepository.Reply(reply)
	if err == nil {
		return err
	}

	return nil
}
