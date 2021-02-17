package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberMemberTagRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationAddLogic {
	return &MemberMemberTagRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationAddLogic) MemberMemberTagRelationAdd(in *ums.MemberMemberTagRelationAddReq) (*ums.MemberMemberTagRelationAddResp, error) {
	_, err := l.svcCtx.UmsMemberMemberTagRelationModel.Insert(
		umsmodel.UmsMemberMemberTagRelation{
			MemberId: in.MemberId,
			TagId:    in.TagId,
		})
	if err != nil {
		return nil, err
	}

	return &ums.MemberMemberTagRelationAddResp{}, nil
}