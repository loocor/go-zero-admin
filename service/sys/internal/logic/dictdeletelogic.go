package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDeleteLogic) DictDelete(in *sys.DictDeleteReq) (*sys.DictDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sys.DictDeleteResp{}, nil
}
