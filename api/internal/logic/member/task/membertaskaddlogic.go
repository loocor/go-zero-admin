package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskAddLogic {
	return MemberTaskAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(req types.MemberTaskAddReq) (*types.MemberTaskAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberTaskAddResp{}, nil
}
