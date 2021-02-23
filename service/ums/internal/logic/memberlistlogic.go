package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberListLogic) MemberList(in *ums.MemberListReq) (*ums.MemberListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberListResp{}, nil
}
