package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryDeleteLogic {
	return GrowthChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(req types.GrowthChangeHistoryDeleteReq) (*types.GrowthChangeHistoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.GrowthChangeHistoryDeleteResp{}, nil
}
