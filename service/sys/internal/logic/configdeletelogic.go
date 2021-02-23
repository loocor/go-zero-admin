package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigDeleteLogic {
	return &ConfigDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigDeleteLogic) ConfigDelete(in *sys.ConfigDeleteReq) (*sys.ConfigDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sys.ConfigDeleteResp{}, nil
}
