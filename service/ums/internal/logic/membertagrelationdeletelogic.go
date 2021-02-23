package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagRelationDeleteLogic {
	return &MemberTagRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagRelationDeleteLogic) MemberTagRelationDelete(in *ums.MemberTagRelationDeleteReq) (*ums.MemberTagRelationDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagRelationDeleteResp{}, nil
}
