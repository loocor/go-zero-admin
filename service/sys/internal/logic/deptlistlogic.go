package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

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
	// todo: add your logic here and delete this line

	return &sys.DeptListResp{}, nil
}
