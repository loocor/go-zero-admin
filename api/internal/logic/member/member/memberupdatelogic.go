package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberUpdateLogic {
	return MemberUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberUpdateLogic) MemberUpdate(req types.MemberUpdateReq) (*types.MemberUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberUpdateResp{}, nil
}
