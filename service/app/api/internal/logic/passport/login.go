package passport

import (
	"context"
	"tinkdance/service/app/api/internal/svc"
	"tinkdance/service/pkg/errorx"
)

type LoginRequest struct{}

type LoginResponse struct{}

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *LoginLogic) Login(req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{}, errorx.WithMsg("我也不知道你干嘛了")
}
