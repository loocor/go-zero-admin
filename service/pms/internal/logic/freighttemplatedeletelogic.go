package logic

import (
	"context"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreightTemplateDeleteLogic {
	return &FreightTemplateDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreightTemplateDeleteLogic) FreightTemplateDelete(in *pms.FreightTemplateDeleteReq) (*pms.FreightTemplateDeleteResp, error) {
	err := l.svcCtx.PmsFreightTemplateModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.FreightTemplateDeleteResp{}, nil
}
