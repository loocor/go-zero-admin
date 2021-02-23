package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryDeleteLogic {
	return IntegrationChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(req types.IntegrationChangeHistoryDeleteReq) (*types.IntegrationChangeHistoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.IntegrationChangeHistoryDeleteResp{}, nil
}
