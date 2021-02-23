package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagRelationListLogic {
	return &MemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagRelationListLogic) MemberTagRelationList(in *ums.MemberTagRelationListReq) (*ums.MemberTagRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagRelationListResp{}, nil
}
