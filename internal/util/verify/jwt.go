package utilverify

import (
	"backend/internal/config"
	modeluser "backend/internal/models/user"
	utillog "backend/internal/util/log"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const path string = "verify/jwt"

var jwtSecret = []byte(config.JWTKEY)

type Claims struct {
	ID       int
	Name     string
	Identity string
	jwt.StandardClaims
}

func GenerateLoginToken(userInfo modeluser.Info) string {
	funcName := "GenerateLoginToken"
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.EXPIRETIME))
	token := ""

	claims := Claims{
		userInfo.Id,
		userInfo.Name,
		userInfo.Identity,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		utillog.FormatString(path, funcName, "生成token出现错误")
		log.Println("WARN: " + err.Error())
	}

	return token
}

func ParseToken(token string) (*Claims, error) {
	funcName := "ParseToken"

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	utillog.FormatString(path, funcName, "验证token出现错误")
	log.Println("WARN: " + err.Error())
	return nil, err
}
