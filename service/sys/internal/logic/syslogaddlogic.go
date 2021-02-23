package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogAddLogic) SysLogAdd(in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	// todo: add your logic here and delete this line

	return &sys.SysLogAddResp{}, nil
}
