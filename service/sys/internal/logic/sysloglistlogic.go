package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogListLogic) SysLogList(in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	// todo: add your logic here and delete this line

	return &sys.SysLogListResp{}, nil
}
