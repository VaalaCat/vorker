package utils

import (
	"errors"
	"time"
	"vorker/conf"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

func init() {
	s := conf.JwtConfig{}
	if err := cleanenv.ReadEnv(&s); err != nil {
		logrus.Panic("load jwt config error")
	}
	conf.JwtConf = &s
}

func SignToken(uid uint) (tokenString string, err error) {
	claim := conf.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.JwtConf.ExpireTime) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UID: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString([]byte(conf.JwtConf.Secret))
	return tokenString, err
}

func ParseToken(tokeStr string) (u *conf.JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(tokeStr, &conf.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtConf.Secret), nil
	})

	if err != nil {
		return nil, errors.New("couldn't handle this token")
	}

	if t, ok := token.Claims.(*conf.JwtClaims); ok && token.Valid {
		return t, nil
	}

	return nil, errors.New("couldn't handle this token")
}
