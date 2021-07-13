package endpoint

import (
	"context"

	errLib "github.com/sapawarga/auth/lib/error"
	"github.com/sapawarga/auth/usecase"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/metadata"
)

// DecodeToken just decode access token
func DecodeToken(s usecase.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errLib.ErrorIncomingContext
		}

		token := md.Get("access_token")
		return s.GetCurrenrLoginFromToken(ctx, token[0])
	}
}

// GetUserDetailByAccessToken get user detail just by parsing active access login
func GetUserDetailByAccessToken(s usecase.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errLib.ErrorIncomingContext
		}
		token := md.Get("access_token")
		user, err := s.GetCurrenrLoginFromToken(ctx, token[0])
		if err != nil {
			return nil, err
		}
		return s.GetAccountDetail(ctx, user.Username)
	}
}
