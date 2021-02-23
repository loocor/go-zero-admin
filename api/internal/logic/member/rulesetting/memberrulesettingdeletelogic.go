package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingDeleteLogic {
	return MemberRuleSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(req types.MemberRuleSettingDeleteReq) (*types.MemberRuleSettingDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberRuleSettingDeleteResp{}, nil
}
