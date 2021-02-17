package logic

import (
	"context"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryUpdateLogic {
	return &ProductAttributeCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryUpdateLogic) ProductAttributeCategoryUpdate(in *pms.ProductAttributeCategoryUpdateReq) (*pms.ProductAttributeCategoryUpdateResp, error) {
	err := l.svcCtx.PmsProductAttributeCategoryModel.Update(
		pmsmodel.PmsProductAttributeCategory{
			Id:             in.Id,
			Name:           in.Name,
			AttributeCount: in.AttributeCount,
			ParamCount:     in.ParamCount,
		})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeCategoryUpdateResp{}, nil
}