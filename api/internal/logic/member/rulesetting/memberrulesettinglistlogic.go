package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingListLogic {
	return MemberRuleSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(req types.MemberRuleSettingListReq) (*types.MemberRuleSettingListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberRuleSettingListResp{}, nil
}
