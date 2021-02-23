package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelAddLogic {
	return &MemberLevelAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(in *ums.MemberLevelAddReq) (*ums.MemberLevelAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLevelAddResp{}, nil
}
