package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagUpdateLogic {
	return MemberTagUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagUpdateLogic) MemberTagUpdate(req types.MemberTagUpdateReq) (*types.MemberTagUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberTagUpdateResp{}, nil
}
