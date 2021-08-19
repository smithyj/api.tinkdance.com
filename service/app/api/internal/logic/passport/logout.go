package passport

import (
	"context"
	"tinkdance/service/app/api/internal/svc"
)

type LogoutRequest struct{}

type LogoutResponse struct{}

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *LogoutLogic) Logout(req *LogoutRequest) (*LogoutResponse, error) {
	return &LogoutResponse{}, nil
}
