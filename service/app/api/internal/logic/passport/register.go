package passport

import (
	"context"
	"tinkdance/service/app/api/internal/svc"
)

type RegisterRequest struct{}

type RegisterResponse struct{}

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *RegisterLogic) Register(req *RegisterRequest) (*RegisterResponse, error) {
	return &RegisterResponse{}, nil
}
