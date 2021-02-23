package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagUpdateLogic {
	return &MemberTagUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagUpdateLogic) MemberTagUpdate(in *ums.MemberTagUpdateReq) (*ums.MemberTagUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagUpdateResp{}, nil
}
