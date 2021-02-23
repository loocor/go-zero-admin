package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressListLogic {
	return &MemberReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressListLogic) MemberReceiveAddressList(in *ums.MemberReceiveAddressListReq) (*ums.MemberReceiveAddressListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberReceiveAddressListResp{}, nil
}
