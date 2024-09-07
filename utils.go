package main

import (
	mathRand "math/rand"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var jwtSecret = []byte("zxcvbnmQWERtyui12345!@#$%^&*()_+ASDFGHJKLE9!3$P2d%YxT4l@Bv&8zQr#Sn2wM6hF5oKdZj7Lp") // Поменяйте на свой секретный ключ
var seededRand = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))                                    // Инициализация seededRand

func CreateToken(userId string, clientIp string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	// Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["client_ip"] = clientIp
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	var err error
	td.AccessToken, err = at.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Refresh Token
	rt := generateRandomString(32)
	td.RefreshToken = rt

	return td, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] // Используем seededRand для генерации случайного индекса
	}
	return string(b)
}
