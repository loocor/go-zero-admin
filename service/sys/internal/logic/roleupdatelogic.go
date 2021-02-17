package logic

import (
	"context"
	"time"

	"go-zero-admin/model/sys"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	_ = l.svcCtx.RoleModel.Update(
		sysmodel.SysRole{
			Id:             in.Id,
			Name:           in.Name,
			Remark:         in.Remark,
			LastUpdateBy:   in.LastUpdateBy,
			LastUpdateTime: time.Now(),
		})

	return &sys.RoleUpdateResp{}, nil
}
