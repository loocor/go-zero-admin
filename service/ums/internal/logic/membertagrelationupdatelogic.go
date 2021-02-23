package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagRelationUpdateLogic {
	return &MemberTagRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagRelationUpdateLogic) MemberTagRelationUpdate(in *ums.MemberTagRelationUpdateReq) (*ums.MemberTagRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagRelationUpdateResp{}, nil
}
