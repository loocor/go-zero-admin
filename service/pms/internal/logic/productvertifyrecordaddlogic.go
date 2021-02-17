package logic

import (
	"context"
	"time"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVerifyRecordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVerifyRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVerifyRecordAddLogic {
	return &ProductVerifyRecordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVerifyRecordAddLogic) ProductVerifyRecordAdd(in *pms.ProductVerifyRecordAddReq) (*pms.ProductVerifyRecordAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsProductVerifyRecordModel.Insert(
		pmsmodel.PmsProductVerifyRecord{
			ProductId:  in.ProductId,
			CreateTime: CreateTime,
			VerifyMan:  in.VerifyMan,
			Status:     in.Status,
			Detail:     in.Detail,
		})
	if err != nil {
		return nil, err
	}

	return &pms.ProductVerifyRecordAddResp{}, nil
}
