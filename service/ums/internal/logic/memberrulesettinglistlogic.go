package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingListLogic {
	return &MemberRuleSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(in *ums.MemberRuleSettingListReq) (*ums.MemberRuleSettingListResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberRuleSettingListResp{}, nil
}
