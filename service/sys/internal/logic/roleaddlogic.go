package logic

import (
	"context"
	"time"

	"go-zero-admin/model/sys"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	_, _ = l.svcCtx.RoleModel.Insert(
		sysmodel.SysRole{
			Name:           in.Name,
			Remark:         in.Remark,
			CreateBy:       in.CreateBy,
			LastUpdateBy:   in.CreateBy,
			LastUpdateTime: time.Now(),
			DelFlag:        0,
		})

	return &sys.RoleAddResp{}, nil
}
