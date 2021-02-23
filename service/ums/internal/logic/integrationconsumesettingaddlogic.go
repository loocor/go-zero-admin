package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingAddLogic {
	return &IntegrationConsumeSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(in *ums.IntegrationConsumeSettingAddReq) (*ums.IntegrationConsumeSettingAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.IntegrationConsumeSettingAddResp{}, nil
}
