package userEntity

import "time"

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type User struct {
	Id              int       `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required"`
	EmailIsVerified bool      `json:"emailIsVerified"`
	Username        string    `json:"username" validate:"required"`
	PhoneNumber     int64     `json:"phoneNumber"`
	IsActive        bool      `json:"isActive"`
	Avatar          string    `json:"avatar"`
	ApiToken        string    `json:"apiToken"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type UserAccess struct {
	Id              int       `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	EmailIsVerified bool      `json:"emailIsVerified"`
	Username        string    `json:"username" validate:"required"`
	PhoneNumber     int64     `json:"phoneNumber"`
	IsActive        bool      `json:"isActive"`
	Avatar          string    `json:"avatar"`
	ApiToken        string    `json:"apiToken"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AccessToken     string    `json:"accessToken"`
}
