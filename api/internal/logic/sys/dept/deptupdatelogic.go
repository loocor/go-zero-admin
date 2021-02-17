package logic

import (
	"context"

	"go-zero-admin/service/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptUpdateLogic {
	return DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptUpdateLogic) DeptUpdate(req types.UpdateDeptReq) (*types.UpdateDeptResp, error) {
	_, err := l.svcCtx.Sys.DeptUpdate(
		l.ctx, &sysclient.DeptUpdateReq{
			Id:           req.Id,
			Name:         req.Name,
			ParentId:     req.ParentId,
			OrderNum:     req.OrderNum,
			LastUpdateBy: "admin", // todo 从 token 里面拿
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.UpdateDeptResp{}, nil
}
