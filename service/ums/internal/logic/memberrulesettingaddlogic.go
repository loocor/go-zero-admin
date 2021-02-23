package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(in *ums.MemberRuleSettingAddReq) (*ums.MemberRuleSettingAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberRuleSettingAddResp{}, nil
}
