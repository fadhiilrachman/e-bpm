package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// hash

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)

	return string(bytes), err
}

func CheckPassword(pwd, dbPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(pwd))

	if err != nil {
		return false
	}

	return true
}

// jwt

type TokenData struct {
	Role		string `json:"id"`
	Username	string `json:"username"`
}

func GetSecretKey() string {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	JWT_SECRET := os.Getenv("JWT_SECRET")
	
	return JWT_SECRET
}

func GenerateToken(role, username string) (string, error) {
	keyString := GetSecretKey()
	key := []byte(keyString)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["role"] = role
	claims["expired_time"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*TokenData, error) {
	keyString := GetSecretKey()
	key := []byte(keyString)

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		role := claims["role"].(string)

		data := &TokenData{
			Role:     role,
			Username: username,
		}

		return data, nil
	} else {
		panic(err)
	}
}

// time

func TimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	t := time.Now().In(loc).Format("2006-01-02 15:04:05")

	return t
}