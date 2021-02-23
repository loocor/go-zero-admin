package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigRoleUpdateLogic {
	return &ConfigRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigRoleUpdateLogic) ConfigRoleUpdate(in *sys.ConfigRoleUpdateReq) (*sys.ConfigRoleUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.ConfigRoleUpdateResp{}, nil
}
