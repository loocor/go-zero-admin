package logic

import (
	"context"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sys.DeptListReq) (*sys.DeptListResp, error) {
	all, err := l.svcCtx.DeptModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
		return nil, err
	}

	var list []*sys.DeptListData
	for _, dept := range *all {

		list = append(list, &sys.DeptListData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dept.DelFlag,
		})
	}

	return &sys.DeptListResp{
		Total: 10,
		List:  list,
	}, nil

}
