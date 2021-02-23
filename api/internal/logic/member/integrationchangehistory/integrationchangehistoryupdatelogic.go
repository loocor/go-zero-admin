package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryUpdateLogic {
	return IntegrationChangeHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(req types.IntegrationChangeHistoryUpdateReq) (*types.IntegrationChangeHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.IntegrationChangeHistoryUpdateResp{}, nil
}
