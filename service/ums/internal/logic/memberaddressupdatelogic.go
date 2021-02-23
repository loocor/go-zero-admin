package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressUpdateLogic {
	return &MemberAddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddressUpdateLogic) MemberAddressUpdate(in *ums.MemberAddressUpdateReq) (*ums.MemberAddressUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberAddressUpdateResp{}, nil
}
