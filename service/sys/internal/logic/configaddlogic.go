package logic

import (
	"context"
	"time"

	sysmodel "go-zero-admin/model/sys"
	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigAddLogic {
	return &ConfigAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigAddLogic) ConfigAdd(in *sys.ConfigAddReq) (*sys.ConfigAddResp, error) {
	_, err := l.svcCtx.ConfigModel.Insert(
		sysmodel.SysConfig{
			Value:          in.Value,
			Label:          in.Label,
			Type:           in.Type,
			Description:    in.Description,
			Sort:           float64(in.Sort),
			CreateBy:       in.CreateBy,
			LastUpdateTime: time.Now(),
			LastUpdateBy:   in.CreateBy,
			Remarks:        in.Remarks,
			DelFlag:        0,
		})
	if err != nil {
		return nil, err
	}

	return &sys.ConfigAddResp{}, nil
}
