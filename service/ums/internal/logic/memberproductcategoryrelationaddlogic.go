package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationAddLogic {
	return &MemberProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationAddLogic) MemberProductCategoryRelationAdd(in *ums.MemberProductCategoryRelationAddReq) (*ums.MemberProductCategoryRelationAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberProductCategoryRelationAddResp{}, nil
}
