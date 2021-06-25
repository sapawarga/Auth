package decoder

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/sapawarga/auth/lib/constant"
	errLib "github.com/sapawarga/auth/lib/error"
	"github.com/sapawarga/auth/model"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) ParsingToken(ctx context.Context, token string) (*model.Actor, error) {
	jwtClaims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_SIGNING_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtClaims.Claims.(jwt.MapClaims)
	if !ok || !jwtClaims.Valid {
		return nil, errLib.ErrorTokenInvalid
	}

	data := claims["data"].(map[string]interface{})

	return &model.Actor{
		Username:    data["username"].(string),
		RoleLabel:   data["roleLabel"].(string),
		LastLoginAt: data["lastLoginAt"],
	}, nil
}
