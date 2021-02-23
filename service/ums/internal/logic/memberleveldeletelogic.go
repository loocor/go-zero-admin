package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelDeleteLogic {
	return &MemberLevelDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(in *ums.MemberLevelDeleteReq) (*ums.MemberLevelDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLevelDeleteResp{}, nil
}
