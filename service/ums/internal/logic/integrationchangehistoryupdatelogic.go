package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryUpdateLogic {
	return &IntegrationChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(in *ums.IntegrationChangeHistoryUpdateReq) (*ums.IntegrationChangeHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.IntegrationChangeHistoryUpdateResp{}, nil
}
