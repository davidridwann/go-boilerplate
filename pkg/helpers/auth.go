package helpers

import (
	"errors"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(id int, code, name, email, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:       id,
		Code:     code,
		Name:     name,
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("token is invalid")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte("")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		log.Err().Fatalln("Invalid JWT Token")
		return nil, err
	}
}

func GetUser(c *gin.Context) *JWTClaim {
	token, err := jwt.ParseWithClaims(strings.Split(c.GetHeader("Authorization"), "Bearer ")[1], &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("supersecretkey"), nil
	})

	if err != nil {
		log.Err().Fatal(err)
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims
	} else {
		log.Err().Fatal(err)
	}

	return nil
}
