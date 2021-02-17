package logic

import (
	"context"

	"go-zero-admin/service/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FreightTemplateUpdateLogic {
	return FreightTemplateUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreightTemplateUpdateLogic) FreightTemplateUpdate(req types.UpdateFreightTemplateReq) (*types.UpdateFreightTemplateResp, error) {
	_, err := l.svcCtx.Pms.FreightTemplateUpdate(
		l.ctx, &pmsclient.FreightTemplateUpdateReq{
			Id:             req.Id,
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

	return &types.UpdateFreightTemplateResp{}, nil
}
