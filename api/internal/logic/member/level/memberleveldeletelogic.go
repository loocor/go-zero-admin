package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelDeleteLogic {
	return MemberLevelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(req types.MemberLevelDeleteReq) (*types.MemberLevelDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberLevelDeleteResp{}, nil
}
