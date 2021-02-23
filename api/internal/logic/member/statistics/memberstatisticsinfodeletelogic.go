package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoDeleteLogic {
	return MemberStatisticsInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(req types.MemberStatisticsInfoDeleteReq) (*types.MemberStatisticsInfoDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberStatisticsInfoDeleteResp{}, nil
}
