package logic

import (
	"context"
	"time"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVerifyRecordUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVerifyRecordUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVerifyRecordUpdateLogic {
	return &ProductVerifyRecordUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVerifyRecordUpdateLogic) ProductVerifyRecordUpdate(in *pms.ProductVerifyRecordUpdateReq) (*pms.ProductVerifyRecordUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsProductVerifyRecordModel.Update(
		pmsmodel.PmsProductVerifyRecord{
			Id:         in.Id,
			ProductId:  in.ProductId,
			CreateTime: CreateTime,
			VerifyMan:  in.VerifyMan,
			Status:     in.Status,
			Detail:     in.Detail,
		})
	if err != nil {
		return nil, err
	}

	return &pms.ProductVerifyRecordUpdateResp{}, nil
}
