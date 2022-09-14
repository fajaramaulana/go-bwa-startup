package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int, levelUser int) (string, error)
}

type jwtService struct {
}

func getSecretKeyJWT() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("JWT_SECRET_KEY")

	return []byte(dsn)
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int, levelUser int) (string, error) {

	payload := jwt.MapClaims{}

	payload["userId"] = userID
	payload["levelUser"] = levelUser
	payload["exp"] = time.Now().Add(time.Hour * 12)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey := getSecretKeyJWT()
	signetToken, err := token.SignedString(secretKey)

	if err != nil {
		return signetToken, err
	}

	return signetToken, nil
}
