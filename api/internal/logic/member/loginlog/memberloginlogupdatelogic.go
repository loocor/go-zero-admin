package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogUpdateLogic {
	return MemberLoginLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(req types.MemberLoginLogUpdateReq) (*types.MemberLoginLogUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberLoginLogUpdateResp{}, nil
}
