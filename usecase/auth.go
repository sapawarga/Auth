package usecase

import (
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/sapawarga/auth/model"
	"github.com/sapawarga/auth/repository"
)

type Auth struct {
	repository repository.AuthI
	actor      string
}

func NewAuth(repo repository.AuthI, actor string) *Auth {
	return &Auth{
		repository: repo,
		actor:      actor,
	}
}

type Actor struct {
	Username  string
	RoleLabel string
}

func (a *Auth) GetActorFromToken(ctx context.Context, token string) (*model.UserDetail, error) {
	actorToken, err := a.parsingToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return a.getActorByUsername(ctx, actorToken.Username)
}

func (a *Auth) parsingToken(ctx context.Context, token string) (*Actor, error) {
	jwtClaims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("someSecretKey"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtClaims.Claims.(jwt.MapClaims)
	if !ok || !jwtClaims.Valid {
		return nil, errors.New("token_not_valid")
	}

	data := claims["data"].(map[string]interface{})

	return &Actor{
		Username:  data["username"].(string),
		RoleLabel: data["roleLabel"].(string),
	}, nil
}

func (a *Auth) getActorByUsername(ctx context.Context, username string) (*model.UserDetail, error) {
	return a.repository.GetActorDetailByUsername(ctx, username)
}
