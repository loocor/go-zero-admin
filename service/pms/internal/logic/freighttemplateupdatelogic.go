package logic

import (
	"context"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreightTemplateUpdateLogic {
	return &FreightTemplateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreightTemplateUpdateLogic) FreightTemplateUpdate(in *pms.FreightTemplateUpdateReq) (*pms.FreightTemplateUpdateResp, error) {
	err := l.svcCtx.PmsFreightTemplateModel.Update(
		pmsmodel.PmsFreightTemplate{
			Id:             in.Id,
			Name:           in.Name,
			ChargeType:     in.ChargeType,
			FirstWeight:    float64(in.FirstWeight),
			FirstFee:       float64(in.FirstFee),
			ContinueWeight: float64(in.ContinueWeight),
			ContinueFee:    float64(in.ContinueFee),
			Dest:           in.Dest,
		})
	if err != nil {
		return nil, err
	}

	return &pms.FreightTemplateUpdateResp{}, nil
}
