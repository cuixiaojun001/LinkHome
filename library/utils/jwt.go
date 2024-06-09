package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

const (
	PhoneRegex           = `^1[3456789]\d{9}$`
	JwtSecret            = "NmUzODk2ZGUtYmZjYy0xMWVjLWI5YTctZjQzMGI5YTUwMzQ2aHVp"
	JwtExpiryHours       = 2
	JwtRefreshExpiryDays = 14
)

// ParseJWTToken 根据token生成用户id
func ParseJWTToken(tokenString string) (int, error) {
	var userId int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat := claims["user_id"].(float64)
		userId = int(userIdFloat)
	} else {
		return -1, fmt.Errorf("token is invalid:%v", err)
	}
	return userId, nil
}
