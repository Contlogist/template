package user

import (
	"fmt"
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

//go:generate mockgen -source=user_token.go -destination=./mock/user_token_mock.go -package=user_mock
type IToken interface {
	Get(ctx *context.Base, userID, companyID int) (*Token, error)
}

type Token struct {
	Access  AccessToken  `json:"access"`
	Refresh RefreshToken `json:"refresh"`
}

type AccessToken struct {
	Token string `json:"token"`
	Hours int    `json:"hours"`
}

type RefreshToken struct {
	Token string `json:"token"`
	Hours int    `json:"hours"`
}

// Get создает модель Tokens.
func (r *Token) Get(ctx *context.Base, uid, cid int) (*Token, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[*Token](ctx, func() (*Token, error) {
		access := AccessToken{}
		access.create(uid, cid)

		refresh := RefreshToken{}
		refresh.create()

		model := Token{
			Access:  access,
			Refresh: refresh,
		}

		return &model, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}

// create создает модель AccessToken.
func (r *AccessToken) create(uid, cid int) {
	hours := 12
	claims := jwt.MapClaims{
		"id":  uid,
		"cid": cid,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	r.Hours = hours
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		logrus.Error("SECRET not env")
	}
	r.Token, _ = token.SignedString([]byte(secret))
}

// create создает модель RefreshToken.
func (r *RefreshToken) create() {
	length := 24
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	charSet := []rune(chars)
	var result string

	for i := 0; i < length; i++ {
		result += string(charSet[rand.Intn(len(charSet))])
	}

	currentTime := time.Now().Format("20060102150405")
	result = fmt.Sprintf("%s%s", currentTime, result[:length-len(currentTime)])

	r.Hours = 24 * 30
	r.Token = result
}
