package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagListLogic {
	return &MemberTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagListLogic) MemberTagList(in *ums.MemberTagListReq) (*ums.MemberTagListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagListResp{}, nil
}
