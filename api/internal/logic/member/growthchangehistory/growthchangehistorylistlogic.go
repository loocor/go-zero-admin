package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryListLogic {
	return GrowthChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(req types.GrowthChangeHistoryListReq) (*types.GrowthChangeHistoryListResp, error) {
	// todo: add your logic here and delete this line

	return &types.GrowthChangeHistoryListResp{}, nil
}
