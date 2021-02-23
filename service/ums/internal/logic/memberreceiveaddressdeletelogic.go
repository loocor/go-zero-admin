package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressDeleteLogic {
	return &MemberReceiveAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressDeleteLogic) MemberReceiveAddressDelete(in *ums.MemberReceiveAddressDeleteReq) (*ums.MemberReceiveAddressDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberReceiveAddressDeleteResp{}, nil
}
