package logic

import (
	"context"

	"go-zero-admin/service/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictUpdateLogic {
	return DictUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictUpdateLogic) DictUpdate(req types.UpdateDictReq) (*types.UpdateDictResp, error) {
	_, err := l.svcCtx.Sys.DictUpdate(
		l.ctx, &sysclient.DictUpdateReq{
			Id:           req.Id,
			Value:        req.Value,
			Label:        req.Label,
			Type:         req.Type,
			Description:  req.Description,
			Sort:         int64(req.Sort),
			Remarks:      req.Remarks,
			LastUpdateBy: "admin", // todo 从 token 里面拿
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.UpdateDictResp{}, nil
}
