package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryUpdateLogic {
	return &GrowthChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(in *ums.GrowthChangeHistoryUpdateReq) (*ums.GrowthChangeHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.GrowthChangeHistoryUpdateResp{}, nil
}
