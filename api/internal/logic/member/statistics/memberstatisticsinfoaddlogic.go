package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoAddLogic {
	return MemberStatisticsInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(req types.MemberStatisticsInfoAddReq) (*types.MemberStatisticsInfoAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberStatisticsInfoAddResp{}, nil
}
