package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigListLogic {
	return ConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigListLogic) ConfigList(req types.ConfigListReq) (*types.ConfigListResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigListResp{}, nil
}
