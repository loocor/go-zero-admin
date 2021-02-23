package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingDeleteLogic {
	return &MemberRuleSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(in *ums.MemberRuleSettingDeleteReq) (*ums.MemberRuleSettingDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberRuleSettingDeleteResp{}, nil
}
