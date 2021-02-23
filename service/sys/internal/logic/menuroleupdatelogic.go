package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRoleUpdateLogic {
	return &MenuRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuRoleUpdateLogic) MenuRoleUpdate(in *sys.MenuRoleUpdateReq) (*sys.MenuRoleUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.MenuRoleUpdateResp{}, nil
}
