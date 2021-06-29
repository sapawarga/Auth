package grpc

import (
	"context"

	"github.com/sapawarga/auth/endpoint"
	"github.com/sapawarga/auth/model"
	"github.com/sapawarga/auth/usecase"
	tpAuth "github.com/sapawarga/proto-file/auth"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

func MakeHandler(fs usecase.Service) tpAuth.AuthHandlerServer {
	decodeTokenHandler := kitgrpc.NewServer(
		endpoint.DecodeToken(fs),
		decodeTokenRequest,
		encodeTokenResponse,
	)

	return &grpcServer{
		decodeTokenHandler,
	}
}

func decodeTokenRequest(ctx context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func encodeTokenResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*model.Actor)

	data := &tpAuth.ResponsesDecodeToken{
		Id:        resp.ID,
		Username:  resp.Username,
		RoleLabel: resp.RoleLabel,
	}

	return data, nil
}
