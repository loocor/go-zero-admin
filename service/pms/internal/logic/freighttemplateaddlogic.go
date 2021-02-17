package logic

import (
	"context"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreightTemplateAddLogic {
	return &FreightTemplateAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreightTemplateAddLogic) FreightTemplateAdd(in *pms.FreightTemplateAddReq) (*pms.FreightTemplateAddResp, error) {
	_, err := l.svcCtx.PmsFreightTemplateModel.Insert(
		pmsmodel.PmsFreightTemplate{
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

	return &pms.FreightTemplateAddResp{}, nil
}
