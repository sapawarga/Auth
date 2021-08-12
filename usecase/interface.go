package usecase

import (
	"context"

	"github.com/sapawarga/auth/model"
)

type Service interface {
	GetCurrenrLoginFromToken(ctx context.Context, token string) (*model.Actor, error)
	GetAccountDetail(ctx context.Context, username string) (*model.UserDetail, error)
}
