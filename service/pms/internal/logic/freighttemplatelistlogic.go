package logic

import (
	"context"
	"fmt"
	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FreightTemplateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreightTemplateListLogic {
	return &FreightTemplateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreightTemplateListLogic) FreightTemplateList(in *pms.FreightTemplateListReq) (*pms.FreightTemplateListResp, error) {
	all, _ := l.svcCtx.PmsFreightTemplateModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.FreightTemplateListData
	for _, item := range *all {

		list = append(list, &pms.FreightTemplateListData{
			Id:             item.Id,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    int64(item.FirstWeight),
			FirstFee:       int64(item.FirstFee),
			ContinueWeight: int64(item.ContinueWeight),
			ContinueFee:    int64(item.ContinueFee),
			Dest:           item.Dest,
		})
	}

	fmt.Println(list)
	return &pms.FreightTemplateListResp{
		Total: 10,
		List:  list,
	}, nil
}
