package repository

import (
	"context"

	"github.com/sapawarga/auth/model"

	_ "github.com/golang/mock/mockgen/model"
)

type AuthI interface {
	GetActorDetailByUsername(ctx context.Context, username string) (*model.UserDetail, error)
	GetActorCurrentLoginByUsername(ctx context.Context, username string) (*model.Actor, error)
}
