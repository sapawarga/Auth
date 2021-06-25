package endpoint

import (
	"context"

	errLib "github.com/sapawarga/auth/lib/error"
	"github.com/sapawarga/auth/usecase"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/metadata"
)

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
