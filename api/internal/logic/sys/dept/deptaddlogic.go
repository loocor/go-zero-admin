package logic

import (
	"context"

	"go-zero-admin/service/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptAddLogic {
	return DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req types.AddDeptReq) (*types.AddDeptResp, error) {
	_, err := l.svcCtx.Sys.DeptAdd(
		l.ctx, &sysclient.DeptAddReq{
			Name:     req.Name,
			ParentId: req.ParentId,
			OrderNum: req.OrderNum,
			CreateBy: "admin", // todo 从 token 里面拿
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.AddDeptResp{}, nil
}
