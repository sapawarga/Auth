package usecase

import (
	"context"

	"github.com/sapawarga/auth/model"
	"github.com/sapawarga/auth/repository"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Auth struct {
	db      repository.AuthI
	decoder repository.JWToken
	actor   string
	logger  kitlog.Logger
}

func NewAuth(repo repository.AuthI, decoder repository.JWToken, actor string, logger kitlog.Logger) *Auth {
	return &Auth{
		db:      repo,
		decoder: decoder,
		actor:   actor,
		logger:  logger,
	}
}

// GetCurrenrLoginFromToken ...
func (a *Auth) GetCurrenrLoginFromToken(ctx context.Context, token string) (*model.Actor, error) {
	logger := kitlog.With(a.logger, "method", "GetCurrentLoginFromToken")
	actorToken, err := a.decoder.ParsingToken(ctx, token)
	if err != nil {
		level.Error(logger).Log("error_parsing_token", err)
		return nil, err
	}

	actor, err := a.db.GetActorCurrentLoginByUsername(ctx, actorToken.Username)
	if err != nil {
		level.Error(logger).Log("error_get_actor", err)
		return nil, err
	}
	actorToken.ID = actor.ID
	return actorToken, nil
}

// GetAccountDetail ...
func (a *Auth) GetAccountDetail(ctx context.Context, username string) (*model.UserDetail, error) {
	logger := kitlog.With(a.logger, "method", "GetAcountDetail")
	user, err := a.db.GetActorDetailByUsername(ctx, username)
	if err != nil {
		level.Error(logger).Log("error_get_actor", err)
		return nil, err
	}
	return user, nil
}
