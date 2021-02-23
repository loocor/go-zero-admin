package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagRelationAddLogic {
	return &MemberTagRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagRelationAddLogic) MemberTagRelationAdd(in *ums.MemberTagRelationAddReq) (*ums.MemberTagRelationAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagRelationAddResp{}, nil
}
