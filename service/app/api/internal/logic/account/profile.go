package account

import (
	"context"
	"tinkdance/service/app/api/internal/svc"
)

type ProfileRequest struct{}

type ProfileResponse struct{}

type ProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *ProfileLogic) Profile(req *ProfileRequest) (*ProfileResponse, error) {
	return &ProfileResponse{}, nil
}
