package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressAddLogic {
	return &MemberReceiveAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressAddLogic) MemberReceiveAddressAdd(in *ums.MemberReceiveAddressAddReq) (*ums.MemberReceiveAddressAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberReceiveAddressAddResp{}, nil
}
