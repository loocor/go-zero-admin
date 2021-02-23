package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryAddLogic {
	return IntegrationChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(req types.IntegrationChangeHistoryAddReq) (*types.IntegrationChangeHistoryAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.IntegrationChangeHistoryAddResp{}, nil
}
