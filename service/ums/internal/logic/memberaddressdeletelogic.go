package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressDeleteLogic {
	return &MemberAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddressDeleteLogic) MemberAddressDelete(in *ums.MemberAddressDeleteReq) (*ums.MemberAddressDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberAddressDeleteResp{}, nil
}
