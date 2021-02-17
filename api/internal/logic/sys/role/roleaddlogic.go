package logic

import (
	"context"

	"go-zero-admin/service/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleAddLogic {
	return RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req types.AddRoleReq) (*types.AddRoleResp, error) {
	_, err := l.svcCtx.Sys.RoleAdd(
		l.ctx, &sysclient.RoleAddReq{
			Name:     req.Name,
			Remark:   req.Remark,
			CreateBy: "admin", // todo 从 token 里面拿
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.AddRoleResp{}, nil
}
