/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 15:17
 * Copyright (c) 2023
 */

package modules

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/oktopriima/gank/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type JwtModules struct {
	SignatureKey    []byte  `json:"signature_key"`
	TokenType       string  `json:"token_type"`
	ExpiredDuration float64 `json:"expired_duration"`
	Issuer          string  `json:"issuer"`
}

type JwtParams struct {
	UserID int
	Email  string
}

type JwtResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiredIn   float64 `json:"expired_in"`
	ExpiredAt   int64   `json:"expired_at"`
}

type JwtExtracted struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	ExpiredAt int64  `json:"expired_at"`
}

func NewJwtModules(env config.Environment) (*JwtModules, error) {
	client := new(JwtModules)
	var err error
	key := env.GetString("jwt.key")
	client.SignatureKey = []byte(key)
	client.Issuer = env.GetString("jwt.issuer")
	client.TokenType = "Bearer"

	client.ExpiredDuration, err = strconv.ParseFloat(env.GetString("jwt.duration"), 64)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (j *JwtModules) GenerateToken(params JwtParams) (*JwtResponse, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredAt := time.Now().Add(time.Second * time.Duration(j.ExpiredDuration))

	myCrypt, err := bcrypt.GenerateFromPassword(j.SignatureKey, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	claims["user_id"] = params.UserID
	claims["email"] = params.Email
	claims["expired_at"] = expiredAt.Unix()
	claims["hash"] = string(myCrypt)

	tokenString, err := token.SignedString(j.SignatureKey)
	if err != nil {
		return nil, err
	}

	resp := new(JwtResponse)
	resp.AccessToken = tokenString
	resp.ExpiredAt = expiredAt.Unix()
	resp.ExpiredIn = j.ExpiredDuration
	resp.TokenType = j.TokenType

	return resp, nil
}

func (j *JwtModules) ExtractToken(token string) (*JwtExtracted, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.SignatureKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		jsonBody, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}

		extract := new(JwtExtracted)
		if err := json.Unmarshal(jsonBody, extract); err != nil {
			return nil, err
		}

		expired := time.Unix(extract.ExpiredAt, 0)
		now := time.Now()

		if now.After(expired) {
			return nil, fmt.Errorf("token expired")
		}

		return extract, nil
	}

	return nil, fmt.Errorf("token not valid")
}
