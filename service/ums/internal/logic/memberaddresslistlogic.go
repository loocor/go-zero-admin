package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressListLogic {
	return &MemberAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddressListLogic) MemberAddressList(in *ums.MemberAddressListReq) (*ums.MemberAddressListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberAddressListResp{}, nil
}
