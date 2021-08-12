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

	getUserDetailHandler := kitgrpc.NewServer(
		endpoint.GetUserDetailByAccessToken(fs),
		decodeTokenRequest,
		encodeUserDetail,
	)

	return &grpcServer{
		decodeTokenHandler,
		getUserDetailHandler,
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

func encodeUserDetail(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*model.UserDetail)

	data := &tpAuth.ResponseUserDetail{
		Id:         resp.ID,
		Username:   resp.Username,
		Email:      resp.Email,
		RoleLabel:  resp.RoleLabel,
		Name:       resp.Name,
		Phone:      resp.Phone,
		PhotoUrl:   resp.PhotoUrl,
		RegencyId:  resp.RegencyID,
		Regency:    resp.Regency,
		BirthDate:  resp.BirthDate.String(),
		DistrictId: resp.DistrictID,
		District:   resp.District,
		VillageId:  resp.VillageID,
		Village:    resp.Village,
		Latitude:   resp.Latitude,
		Longitude:  resp.Longitude,
	}

	return data, nil
}
