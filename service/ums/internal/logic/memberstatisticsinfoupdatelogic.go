package logic

import (
	"context"

	"zdmin/service/ums/internal/svc"
	"zdmin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoUpdateLogic {
	return &MemberStatisticsInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(in *ums.MemberStatisticsInfoUpdateReq) (*ums.MemberStatisticsInfoUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberStatisticsInfoUpdateResp{}, nil
}
