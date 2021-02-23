package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleUpdateLogic {
	return &UserRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleUpdateLogic) UserRoleUpdate(in *sys.UserRoleUpdateReq) (*sys.UserRoleUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserRoleUpdateResp{}, nil
}
