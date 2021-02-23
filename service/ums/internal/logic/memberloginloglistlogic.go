package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(in *ums.MemberLoginLogListReq) (*ums.MemberLoginLogListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLoginLogListResp{}, nil
}
