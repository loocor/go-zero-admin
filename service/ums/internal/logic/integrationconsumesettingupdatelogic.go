package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

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
	err := l.svcCtx.UmsIntegrationConsumeSettingModel.Update(
		umsmodel.UmsIntegrationConsumeSetting{
			Id:                 in.Id,
			DeductionPerAmount: in.DeductionPerAmount,
			MaxPercentPerOrder: in.MaxPercentPerOrder,
			UseUnit:            in.UseUnit,
			CouponStatus:       in.CouponStatus,
		})
	if err != nil {
		return nil, err
	}

	return &ums.IntegrationConsumeSettingUpdateResp{}, nil
}
