package logic

import (
	"context"
	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	l.svcCtx.RoleModel.Delete(in.Id)

	return &sys.RoleDeleteResp{}, nil
}
