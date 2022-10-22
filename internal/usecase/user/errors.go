package userUseCase

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrUnexpected = errors.New("Unexpected Error")
var ErrEmailExists = errors.New("Email already exists")
var ErrPasswordFailed = errors.New("Email atau Password Salah")
