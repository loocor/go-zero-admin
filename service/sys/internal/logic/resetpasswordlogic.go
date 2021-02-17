package logic

import (
	"context"
	"time"

	"go-zero-admin/model/sys"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(in *sys.ReSetPasswordReq) (*sys.ReSetPasswordResp, error) {

	_ = l.svcCtx.UserModel.Update(
		sysmodel.SysUser{
			Id:             in.Id,
			Password:       "123456",
			Salt:           "123456",
			LastUpdateBy:   in.LastUpdateBy,
			LastUpdateTime: time.Now(),
		})

	return &sys.ReSetPasswordResp{}, nil
}
