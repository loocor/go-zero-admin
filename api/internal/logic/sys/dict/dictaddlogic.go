package logic

import (
	"context"
	"strconv"

	"go-zero-admin/service/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

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

func (l *DictAddLogic) DictAdd(req types.AddDictReq) (*types.AddDictResp, error) {
	sort, _ := strconv.ParseInt(req.Sort, 10, 64)
	_, err := l.svcCtx.Sys.DictAdd(
		l.ctx, &sysclient.DictAddReq{
			Value:       req.Value,
			Label:       req.Label,
			Type:        req.Type,
			Description: req.Description,
			Sort:        sort,
			Remarks:     req.Remarks,
			CreateBy:    "admin", // todo 从 token 里面拿
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.AddDictResp{}, nil
}
