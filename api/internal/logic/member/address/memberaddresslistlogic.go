package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressListLogic {
	return MemberAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressListLogic) MemberAddressList(req types.MemberAddressListReq) (*types.MemberAddressListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberAddressListResp{}, nil
}
