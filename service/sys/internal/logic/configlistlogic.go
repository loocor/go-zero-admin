package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigListLogic {
	return &ConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigListLogic) ConfigList(in *sys.ConfigListReq) (*sys.ConfigListResp, error) {
	// todo: add your logic here and delete this line

	return &sys.ConfigListResp{}, nil
}
