package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingDeleteLogic {
	return IntegrationConsumeSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(req types.IntegrationConsumeSettingDeleteReq) (*types.IntegrationConsumeSettingDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.IntegrationConsumeSettingDeleteResp{}, nil
}
