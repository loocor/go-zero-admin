package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingDeleteLogic {
	return &IntegrationConsumeSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(in *ums.IntegrationConsumeSettingDeleteReq) (*ums.IntegrationConsumeSettingDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.IntegrationConsumeSettingDeleteResp{}, nil
}
