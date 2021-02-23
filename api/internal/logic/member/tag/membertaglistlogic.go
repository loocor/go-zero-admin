package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagListLogic {
	return MemberTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagListLogic) MemberTagList(req types.MemberTagListReq) (*types.MemberTagListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberTagListResp{}, nil
}
