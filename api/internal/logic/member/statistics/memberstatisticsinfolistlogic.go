package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoListLogic {
	return MemberStatisticsInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoListLogic) MemberStatisticsInfoList(req types.MemberStatisticsInfoListReq) (*types.MemberStatisticsInfoListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberStatisticsInfoListResp{}, nil
}
