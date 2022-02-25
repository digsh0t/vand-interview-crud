package authentication

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wintltr/vand-interview-crud-project/util"
)

type TokenData struct {
	Username   string
	Role       string
	Userid     int
	Exp        uint64
	Twofa      string
	Authorized bool
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userId
	claims["exp"] = time.Now().Add(time.Hour + time.Hour/2).Unix() // Token expires after 1 hour and 30 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Activate env file
	util.EnvInit()
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}