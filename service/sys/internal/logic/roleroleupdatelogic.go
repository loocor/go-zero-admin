package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRoleUpdateLogic {
	return &RoleRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleRoleUpdateLogic) RoleRoleUpdate(in *sys.RoleRoleUpdateReq) (*sys.RoleRoleUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.RoleRoleUpdateResp{}, nil
}
