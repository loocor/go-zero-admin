package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationListLogic {
	return &MemberProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationListLogic) MemberProductCategoryRelationList(in *ums.MemberProductCategoryRelationListReq) (*ums.MemberProductCategoryRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberProductCategoryRelationListResp{}, nil
}
