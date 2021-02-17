package logic

import (
	"context"
	"fmt"

	"go-zero-admin/service/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FreightTemplateListLogic {
	return FreightTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreightTemplateListLogic) FreightTemplateList(req types.ListFreightTemplateReq) (*types.ListFreightTemplateResp, error) {
	resp, err := l.svcCtx.Pms.FreightTemplateList(
		l.ctx, &pmsclient.FreightTemplateListReq{
			Current:  req.Current,
			PageSize: req.PageSize,
		},
	)

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtFreightTemplateData

	for _, item := range resp.List {
		list = append(
			list, &types.ListtFreightTemplateData{
				Id:             item.Id,
				Name:           item.Name,
				ChargeType:     item.ChargeType,
				FirstWeight:    float64(item.FirstWeight),
				FirstFee:       float64(item.FirstFee),
				ContinueWeight: float64(item.ContinueWeight),
				ContinueFee:    float64(item.FirstFee),
				Dest:           item.Dest,
			},
		)
	}

	return &types.ListFreightTemplateResp{
		Current:  req.Current,
		Data:     nil,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
