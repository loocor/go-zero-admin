package logic

import (
	"context"

	"go-zero-admin/service/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FreightTemplateAddLogic {
	return FreightTemplateAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreightTemplateAddLogic) FreightTemplateAdd(req types.AddFreightTemplateReq) (*types.AddFreightTemplateResp, error) {
	_, err := l.svcCtx.Pms.FreightTemplateAdd(
		l.ctx, &pmsclient.FreightTemplateAddReq{
			Name:           req.Name,
			ChargeType:     req.ChargeType,
			FirstWeight:    int64(req.FirstWeight),
			FirstFee:       int64(req.FirstFee),
			ContinueWeight: int64(req.ContinueWeight),
			ContinueFee:    int64(req.FirstFee),
			Dest:           req.Dest,
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.AddFreightTemplateResp{}, nil
}
