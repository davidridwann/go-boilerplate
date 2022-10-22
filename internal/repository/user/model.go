package userRepository

import (
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID              uint   `gorm:"primarykey"`
	Code            string `gorm:"unique"`
	Name            string
	Email           string `gorm:"unique"`
	EmailIsVerified bool
	Username        string `gorm:"unique, omitempty"`
	Password        string
	PhoneNumber     int64 `gorm:"omitempty"`
	IsActive        bool
	Avatar          string `gorm:"omitempty"`
	ApiToken        string `gorm:"index, omitempty"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (user *User) ToEntity() *userEntity.User {
	return &userEntity.User{
		Id:              int(user.ID),
		Code:            user.Code,
		Name:            user.Name,
		Email:           user.Email,
		EmailIsVerified: user.EmailIsVerified,
		Username:        user.Username,
		Password:        user.Password,
		PhoneNumber:     user.PhoneNumber,
		IsActive:        user.IsActive,
		Avatar:          user.Avatar,
		ApiToken:        user.ApiToken,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}

func (User) FromEntity(user *userEntity.User) *User {
	return &User{
		ID:              uint(user.Id),
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Code:            user.Code,
		Name:            user.Name,
		Email:           user.Email,
		EmailIsVerified: user.EmailIsVerified,
		Username:        user.Username,
		Password:        user.Password,
		PhoneNumber:     user.PhoneNumber,
		IsActive:        user.IsActive,
		Avatar:          user.Avatar,
		ApiToken:        user.ApiToken,
	}
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
