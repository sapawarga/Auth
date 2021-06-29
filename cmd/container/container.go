package container

import (
	"github.com/sapawarga/auth/repository/decoder"
	"github.com/sapawarga/auth/repository/mysql"
	"github.com/sapawarga/auth/usecase"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	AuthService usecase.Service
}

func Init(conn *sqlx.DB, logger log.Logger, actor string) *Container {
	return &Container{
		AuthService: usecase.NewAuth(mysql.NewAuth(conn), decoder.NewToken(), actor, logger),
	}
}
