package usecase

import (
	"context"

	"github.com/sapawarga/auth/model"
)

type Service interface {
	GetActorFromToken(ctx context.Context, token string) (*model.UserDetail, error)
}
