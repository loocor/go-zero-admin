package logic

import (
	"context"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

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
	err := l.svcCtx.UmsIntegrationConsumeSettingModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.IntegrationConsumeSettingDeleteResp{}, nil
}
