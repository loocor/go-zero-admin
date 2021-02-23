package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigAddLogic {
	return ConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigAddLogic) ConfigAdd(req types.ConfigAddReq) (*types.ConfigAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigAddResp{}, nil
}
