package usecase

import (
	"context"

	"github.com/sapawarga/auth/model"
)

type Service interface {
	GetCurrenrLoginFromToken(ctx context.Context, token string) (*model.Actor, error)
}
