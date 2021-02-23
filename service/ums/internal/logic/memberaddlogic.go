package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddLogic) MemberAdd(in *ums.MemberAddReq) (*ums.MemberAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberAddResp{}, nil
}
