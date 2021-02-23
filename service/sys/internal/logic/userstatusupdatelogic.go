package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserStatusUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserStatusUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserStatusUpdateLogic {
	return &UserStatusUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserStatusUpdateLogic) UserStatusUpdate(in *sys.UserStatusReq) (*sys.UserStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserStatusResp{}, nil
}
