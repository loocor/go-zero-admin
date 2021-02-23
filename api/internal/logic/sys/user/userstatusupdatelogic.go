package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserStatusUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserStatusUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserStatusUpdateLogic {
	return UserStatusUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserStatusUpdateLogic) UserStatusUpdate(req types.UserStatusUpdateReq) (*types.UserStatusUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserStatusUpdateResp{}, nil
}
