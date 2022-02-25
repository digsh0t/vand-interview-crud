package authentication

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wintltr/vand-interview-crud-project/util"
)

type TokenData struct {
	Userid     int
	Exp        uint64

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

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*TokenData, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId, err := strconv.Atoi(fmt.Sprintf("%.f", claims["userid"]))
		if err != nil {
			return nil, err
		}
		exp, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &TokenData{
			Userid:     userId,
			Exp:        exp,
		}, nil
	}
	return nil, err
}