package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagDeleteLogic {
	return MemberTagDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(req types.MemberTagDeleteReq) (*types.MemberTagDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberTagDeleteResp{}, nil
}
