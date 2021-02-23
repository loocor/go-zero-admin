package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictAddLogic {
	return DictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictAddLogic) DictAdd(req types.DictAddReq) (*types.DictAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.DictAddResp{}, nil
}
