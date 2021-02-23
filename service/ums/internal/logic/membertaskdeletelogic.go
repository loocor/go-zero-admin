package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskDeleteLogic {
	return &MemberTaskDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(in *ums.MemberTaskDeleteReq) (*ums.MemberTaskDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTaskDeleteResp{}, nil
}
