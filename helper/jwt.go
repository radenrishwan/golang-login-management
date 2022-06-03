package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"login-management/model/domain"
	"os"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func ClaimToken(token string) (domain.Session, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method : %s", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return domain.Session{}, errors.New("invalid token")
	}

	// claim
	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		return domain.Session{
			Id:        claims["id"].(string),
			UserId:    claims["user_id"].(string),
			CreatedAt: claims["created_at"].(string),
			ExpiredAt: claims["expired_at"].(string),
		}, nil
	}

	return domain.Session{}, errors.New("invalid token")
}

func sessionToMapClaims(session *domain.Session) jwt.MapClaims {
	return jwt.MapClaims{
		"id":         session.Id,
		"user_id":    session.UserId,
		"created_at": session.CreatedAt,
		"expired_at": session.ExpiredAt,
	}
}

func NewTokenString(session domain.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionToMapClaims(&session))

	return token.SignedString(secretKey)
}
