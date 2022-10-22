package userUseCase

import userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"

type IUseCase interface {
	Get(id int) (*userEntity.User, error)
	Register(in userEntity.User) (*userEntity.User, error)
	Login(string, string) (*userEntity.UserAccess, error)
}
