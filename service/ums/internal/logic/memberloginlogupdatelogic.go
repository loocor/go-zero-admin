package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogUpdateLogic {
	return &MemberLoginLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(in *ums.MemberLoginLogUpdateReq) (*ums.MemberLoginLogUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLoginLogUpdateResp{}, nil
}
