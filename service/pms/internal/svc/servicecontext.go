package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"go-zero-admin/model/pms"
	"go-zero-admin/service/pms/internal/config"
)

type ServiceContext struct {
	c config.Config

	PmsAlbumModel                            *pmsmodel.PmsAlbumModel
	PmsAlbumPicModel                         *pmsmodel.PmsAlbumPicModel
	PmsBrandModel                            *pmsmodel.PmsBrandModel
	PmsCommentModel                          *pmsmodel.PmsCommentModel
	PmsCommentReplayModel                    *pmsmodel.PmsCommentReplayModel
	PmsFreightTemplateModel                  *pmsmodel.PmsFreightTemplateModel
	PmsMemberPriceModel                      *pmsmodel.PmsMemberPriceModel
	PmsProductAttributeCategoryModel         *pmsmodel.PmsProductAttributeCategoryModel
	PmsProductAttributeModel                 *pmsmodel.PmsProductAttributeModel
	PmsProductAttributeValueModel            *pmsmodel.PmsProductAttributeValueModel
	PmsProductCategoryAttributeRelationModel *pmsmodel.PmsProductCategoryAttributeRelationModel
	PmsProductCategoryModel                  *pmsmodel.PmsProductCategoryModel
	PmsProductFullReductionModel             *pmsmodel.PmsProductFullReductionModel
	PmsProductLadderModel                    *pmsmodel.PmsProductLadderModel
	PmsProductModel                          *pmsmodel.PmsProductModel
	PmsProductOperateLogModel                *pmsmodel.PmsProductOperateLogModel
	PmsProductVerifyRecordModel              *pmsmodel.PmsProductVerifyRecordModel
	PmsSkuStockModel                         *pmsmodel.PmsSkuStockModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		PmsAlbumModel:                            pmsmodel.NewPmsAlbumModel(sqlConn),
		PmsAlbumPicModel:                         pmsmodel.NewPmsAlbumPicModel(sqlConn),
		PmsBrandModel:                            pmsmodel.NewPmsBrandModel(sqlConn),
		PmsCommentModel:                          pmsmodel.NewPmsCommentModel(sqlConn),
		PmsCommentReplayModel:                    pmsmodel.NewPmsCommentReplayModel(sqlConn),
		PmsFreightTemplateModel:                  pmsmodel.NewPmsFreightTemplateModel(sqlConn),
		PmsMemberPriceModel:                      pmsmodel.NewPmsMemberPriceModel(sqlConn),
		PmsProductAttributeCategoryModel:         pmsmodel.NewPmsProductAttributeCategoryModel(sqlConn),
		PmsProductAttributeModel:                 pmsmodel.NewPmsProductAttributeModel(sqlConn),
		PmsProductAttributeValueModel:            pmsmodel.NewPmsProductAttributeValueModel(sqlConn),
		PmsProductCategoryAttributeRelationModel: pmsmodel.NewPmsProductCategoryAttributeRelationModel(sqlConn),
		PmsProductCategoryModel:                  pmsmodel.NewPmsProductCategoryModel(sqlConn),
		PmsProductFullReductionModel:             pmsmodel.NewPmsProductFullReductionModel(sqlConn),
		PmsProductLadderModel:                    pmsmodel.NewPmsProductLadderModel(sqlConn),
		PmsProductModel:                          pmsmodel.NewPmsProductModel(sqlConn),
		PmsProductOperateLogModel:                pmsmodel.NewPmsProductOperateLogModel(sqlConn),
		PmsProductVerifyRecordModel:              pmsmodel.NewPmsProductVerifyRecordModel(sqlConn),
		PmsSkuStockModel:                         pmsmodel.NewPmsSkuStockModel(sqlConn),
	}
}
