package logic

import (
	"context"

	"go-zero-admin/service/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FreightTemplateDeleteLogic {
	return FreightTemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreightTemplateDeleteLogic) FreightTemplateDelete(req types.DeleteFreightTemplateReq) (*types.DeleteFreightTemplateResp, error) {
	_, _ = l.svcCtx.Pms.FreightTemplateDelete(
		l.ctx, &pmsclient.FreightTemplateDeleteReq{
			Id: req.Id,
		},
	)

	return &types.DeleteFreightTemplateResp{}, nil
}
