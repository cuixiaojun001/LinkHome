package test

//package main
//
//import (
//	"errors"
//	"fmt"
//	"github.com/golang-jwt/jwt"
//	"time"
//)
//
//var JwtSecret = []byte("NmUzODk2ZGUtYmZjYy0xMWVjLWI5YTctZjQzMGI5YTUwMzQ2aHVp")
//
//type Claims struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}
//
//func generateJWT(username string) (string, error) {
//	expirationTime := time.Now().Add(5 * time.Minute)
//	claims := &Claims{
//		Username: username,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString(JwtSecret)
//}
//
//func parseJWT(tokenStr string) (*Claims, error) {
//	claims := &Claims{}
//
//	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//		return JwtSecret, nil
//	})
//
//	if err != nil {
//		if errors.Is(err, jwt.ErrSignatureInvalid) {
//			return nil, fmt.Errorf("invalid signature")
//		}
//		return nil, fmt.Errorf("could not parse token")
//	}
//
//	if !token.Valid {
//		return nil, fmt.Errorf("invalid token")
//	}
//
//	return claims, nil
//}
//
//func main() {
//	// Generate a new token
//	token, err := generateJWT("cuixiaojun")
//	if err != nil {
//		fmt.Println("Error generating token:", err)
//		return
//	}
//
//	fmt.Println("Generated token:", token)
//
//	// Parse and validate the token
//	claims, err := parseJWT(token)
//	if err != nil {
//		fmt.Println("Error parsing token:", err)
//		return
//	}
//
//	fmt.Println("Token is valid. Username:", claims.Username)
//}
