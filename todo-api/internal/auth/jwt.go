package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string) (string, error) {
	// Implementation for generating a JWT
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// In a real application, use a secure key from config/env
	secretKey := []byte("SecretKey")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(token string) (int, error) {
	// Implementation for validating a JWT
	fmt.Println("Validating Token:", token)
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		fmt.Printf("Token Method: %v\n", token.Method)
		return []byte("SecretKey"), nil
	})
	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return 0, err
	}
	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok || !tokenObj.Valid {
		fmt.Printf("isok : %v tokenObj.Valid: %v\n", ok, tokenObj.Valid)
		return 0, fmt.Errorf("invalid token")
	}
	// userID, ok := claims["sub"].(float64)
	// if !ok {
	// 	fmt.Printf("Error asserting userID type: %T\n", claims["sub"])
	// 	return "", fmt.Errorf("invalid token claims")
	// }
	userID := 0
	switch v := claims["sub"].(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	case string:
		var err error
		userID, err = strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("invalid token claims")
		}
	default:
		return 0, fmt.Errorf("invalid token claims")
	}
	fmt.Printf("Extracted UserID: %v\n", userID)
	return userID, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	// "Bearer <token>"
	if len(bearerToken) > 7 && strings.ToLower(bearerToken[:7]) == "bearer " {
		token := bearerToken[7:]
		fmt.Println("Extracted Token:", token)
		return token
	}
	return ""
}
