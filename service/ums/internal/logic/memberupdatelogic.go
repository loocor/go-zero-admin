package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdateLogic {
	return &MemberUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberUpdateLogic) MemberUpdate(in *ums.MemberUpdateReq) (*ums.MemberUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberUpdateResp{}, nil
}
