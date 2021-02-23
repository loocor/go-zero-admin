package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskDeleteLogic {
	return MemberTaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(req types.MemberTaskDeleteReq) (*types.MemberTaskDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberTaskDeleteResp{}, nil
}
