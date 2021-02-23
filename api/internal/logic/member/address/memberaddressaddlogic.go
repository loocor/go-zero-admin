package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressAddLogic {
	return MemberAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(req types.MemberAddressAddReq) (*types.MemberAddressAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberAddressAddResp{}, nil
}
