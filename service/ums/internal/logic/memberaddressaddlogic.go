package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressAddLogic {
	return &MemberAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(in *ums.MemberAddressAddReq) (*ums.MemberAddressAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberAddressAddResp{}, nil
}
