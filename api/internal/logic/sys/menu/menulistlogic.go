package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuListLogic {
	return MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req types.MenuListReq) (*types.MenuListResp, error) {
	// todo: add your logic here and delete this line

	return &types.MenuListResp{}, nil
}
