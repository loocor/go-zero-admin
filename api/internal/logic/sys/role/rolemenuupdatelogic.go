package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleMenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleMenuUpdateLogic {
	return RoleMenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleMenuUpdateLogic) RoleMenuUpdate(req types.RoleMenuUpdateReq) (*types.RoleMenuUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.RoleMenuUpdateResp{}, nil
}
