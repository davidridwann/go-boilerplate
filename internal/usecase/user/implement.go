package userUseCase

import (
	"errors"
	"fmt"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	userRepository "github.com/davidridwann/wlb-test.git/internal/repository/user"
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type IUseCaseImplementation struct {
	userRepository userRepository.UserRepository
}

func NewUseCase(userRepository userRepository.UserRepository) IUseCase {
	return &IUseCaseImplementation{userRepository}
}

func (uc *IUseCaseImplementation) Get(id int) (*userEntity.User, error) {
	data, err := uc.userRepository.Get(id)
	if err != nil {
		if errors.Is(err, userRepository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Register(in userEntity.User) (*userEntity.User, error) {
	_, err := uc.userRepository.GetByEmail(in.Email)
	if err == nil {
		return nil, ErrEmailExists
	}

	fmt.Println(in)
	data, err := uc.userRepository.Create(in)
	if err != nil {
		return nil, ErrUnexpected
	}

	return data, nil
}

func (uc *IUseCaseImplementation) Login(email string, password string) (*userEntity.UserAccess, error) {
	user, err := uc.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, userRepository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrUnexpected
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrPasswordFailed
	}

	accessToken, err := helpers.GenerateJWT(user.Id, user.Code, user.Name, user.Email, user.Username)
	if err != nil {
		return nil, errors.New("JWT Error Create")
	}

	return &userEntity.UserAccess{
		Id:              user.Id,
		Code:            user.Code,
		Name:            user.Name,
		Email:           user.Email,
		EmailIsVerified: user.EmailIsVerified,
		Username:        user.Username,
		Phone:           user.Phone,
		IsActive:        user.IsActive,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		AccessToken:     accessToken,
	}, nil
}
