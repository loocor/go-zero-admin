package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingUpdateLogic {
	return MemberRuleSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(req types.MemberRuleSettingUpdateReq) (*types.MemberRuleSettingUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberRuleSettingUpdateResp{}, nil
}
