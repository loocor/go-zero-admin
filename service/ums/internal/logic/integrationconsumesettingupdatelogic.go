package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingUpdateLogic {
	return &IntegrationConsumeSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(in *ums.IntegrationConsumeSettingUpdateReq) (*ums.IntegrationConsumeSettingUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.IntegrationConsumeSettingUpdateResp{}, nil
}
