package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelListLogic {
	return &MemberLevelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelListLogic) MemberLevelList(in *ums.MemberLevelListReq) (*ums.MemberLevelListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLevelListResp{}, nil
}
