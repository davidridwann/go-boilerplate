package userRepository

import (
	"errors"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var ErrUserNotFound = errors.New("User not found")
var ErrUnexpected = errors.New("Unexpected Error")
var ErrEmailExists = errors.New("Email already exists")

type UserRepository interface {
	Get(id int) (*userEntity.User, error)
	GetByEmail(string) (*userEntity.User, error)
	Create(in userEntity.User) (*userEntity.User, error)
	Login(string, string) (*userEntity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return &Repository{db}
}

func (r *Repository) Get(id int) (*userEntity.User, error) {
	userData := &User{}

	err := r.db.Raw(`SELECT * FROM users
		WHERE id = ?
	`, id).First(&userData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	data := userData.ToEntity()
	return data, err
}

func (r *Repository) GetByEmail(email string) (*userEntity.User, error) {
	userData := &User{}

	err := r.db.Raw(`SELECT * FROM users
		WHERE email = ?
	`, email).First(&userData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	data := userData.ToEntity()
	return data, err
}

func (r *Repository) Create(in userEntity.User) (*userEntity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 14)
	if err != nil {
		return nil, err
	}

	code := uuid.New()
	data := User{
		Code:            code.String(),
		Name:            in.Name,
		Email:           in.Email,
		EmailIsVerified: false,
		Username:        in.Username,
		Password:        string(hashedPassword),
		PhoneNumber:     0,
		IsActive:        false,
		Avatar:          "",
		ApiToken:        "",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}

	err = r.db.Table("users").Create(&data).Error
	if err != nil {
		return nil, err
	}

	user, _ := r.GetByEmail(in.Email)
	return user, nil
}

func (r *Repository) Login(email, password string) (*userEntity.User, error) {
	_, err := r.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return nil, err
}
