package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogAddLogic {
	return MemberLoginLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(req types.MemberLoginLogAddReq) (*types.MemberLoginLogAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberLoginLogAddResp{}, nil
}
