package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberListLogic {
	return MemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberListLogic) MemberList(req types.MemberListReq) (*types.MemberListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberListResp{}, nil
}
