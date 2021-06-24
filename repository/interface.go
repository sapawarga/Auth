package repository

import (
	"context"

	"github.com/sapawarga/auth/model"
)

type AuthI interface {
	GetActorDetailByUsername(ctx context.Context, username string) (*model.UserDetail, error)
}
