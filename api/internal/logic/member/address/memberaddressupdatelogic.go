package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressUpdateLogic {
	return MemberAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressUpdateLogic) MemberAddressUpdate(req types.MemberAddressUpdateReq) (*types.MemberAddressUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberAddressUpdateResp{}, nil
}
