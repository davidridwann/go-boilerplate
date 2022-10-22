package userEntity

import "time"

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type User struct {
	Id              int       `json:"id" swaggerignore:"true"`
	Code            string    `json:"code" swaggerignore:"true"`
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"required,min=6"`
	EmailIsVerified bool      `json:"emailIsVerified" swaggerignore:"true"`
	Username        string    `json:"username" binding:"required"`
	Phone           string    `json:"phone"`
	IsActive        bool      `json:"isActive" swaggerignore:"true"`
	CreatedAt       time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt       time.Time `json:"updated_at" swaggerignore:"true"`
	DeletedAt       time.Time `json:"deleted_at" swaggerignore:"true"`
}

type UserAccess struct {
	Id              int       `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	EmailIsVerified bool      `json:"emailIsVerified"`
	Username        string    `json:"username" binding:"required"`
	Phone           string    `json:"phone"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AccessToken     string    `json:"accessToken"`
}
