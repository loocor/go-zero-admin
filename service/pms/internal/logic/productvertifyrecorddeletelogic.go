package logic

import (
	"context"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVerifyRecordDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVerifyRecordDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVerifyRecordDeleteLogic {
	return &ProductVerifyRecordDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVerifyRecordDeleteLogic) ProductVerifyRecordDelete(in *pms.ProductVerifyRecordDeleteReq) (*pms.ProductVerifyRecordDeleteResp, error) {
	err := l.svcCtx.PmsProductVerifyRecordModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductVerifyRecordDeleteResp{}, nil
}
