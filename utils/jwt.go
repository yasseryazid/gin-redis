package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yasseryazid/technical-test/config"
)

func GenerateJWT(userID uint, username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET not set in environment")
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix() // Berlaku 24 jam

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	err = config.RedisClient.Set(ctx, tokenString, userID, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET not set in environment")
	}

	ctx := context.Background()
	_, err := config.RedisClient.Get(ctx, tokenString).Result()
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func LogoutJWT(tokenString string) error {
	ctx := context.Background()
	err := config.RedisClient.Del(ctx, tokenString).Err()
	return err
}
