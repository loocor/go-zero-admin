package logic

import (
	"context"
	"time"

	"go-zero-admin/model/sys"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(in *sys.UserStatusReq) (*sys.UserStatusResp, error) {
	_ = l.svcCtx.UserModel.Update(
		sysmodel.SysUser{
			Id:             in.Id,
			Status:         in.Status,
			LastUpdateBy:   in.LastUpdateBy,
			LastUpdateTime: time.Now(),
		})

	return &sys.UserStatusResp{}, nil
}
