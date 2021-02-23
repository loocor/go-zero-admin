package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskListLogic {
	return &MemberTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskListLogic) MemberTaskList(in *ums.MemberTaskListReq) (*ums.MemberTaskListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTaskListResp{}, nil
}
