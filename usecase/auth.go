package usecase

import (
	"context"

	"github.com/sapawarga/auth/model"
	"github.com/sapawarga/auth/repository"
)

type Auth struct {
	db      repository.AuthI
	decoder repository.JWToken
	actor   string
}

func NewAuth(repo repository.AuthI, decoder repository.JWToken, actor string) *Auth {
	return &Auth{
		db:      repo,
		decoder: decoder,
		actor:   actor,
	}
}

func (a *Auth) GetCurrenrLoginFromToken(ctx context.Context, token string) (*model.Actor, error) {
	actorToken, err := a.decoder.ParsingToken(ctx, token)
	if err != nil {
		return nil, err
	}

	actor, err := a.db.GetActorCurrentLoginByUsername(ctx, actorToken.Username)
	if err != nil {
		return nil, err
	}
	actorToken.ID = actor.ID
	return actorToken, nil
}

// func (a *Auth) getActorByUsername(ctx context.Context, username string) (*model.UserDetail, error) {
// 	return a.repository.GetActorDetailByUsername(ctx, username)
// }
