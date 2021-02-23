package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *sys.DictUpdateReq) (*sys.DictUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.DictUpdateResp{}, nil
}
