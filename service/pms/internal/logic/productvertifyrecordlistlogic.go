package logic

import (
	"context"
	"fmt"
	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVerifyRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVerifyRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVerifyRecordListLogic {
	return &ProductVerifyRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVerifyRecordListLogic) ProductVerifyRecordList(in *pms.ProductVerifyRecordListReq) (*pms.ProductVerifyRecordListResp, error) {
	all, _ := l.svcCtx.PmsProductVerifyRecordModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductVerifyRecordListData
	for _, item := range *all {

		list = append(list, &pms.ProductVerifyRecordListData{
			Id:         item.Id,
			ProductId:  item.ProductId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			VerifyMan:  item.VerifyMan,
			Status:     item.Status,
			Detail:     item.Detail,
		})
	}

	fmt.Println(list)
	return &pms.ProductVerifyRecordListResp{
		Total: 10,
		List:  list,
	}, nil
}
