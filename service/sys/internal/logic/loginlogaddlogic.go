package logic

import (
	"context"
	"time"

	"go-zero-admin/model/sys"

	"go-zero-admin/service/sys/internal/svc"
	"go-zero-admin/service/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(
		sysmodel.SysLoginLog{
			UserName:       in.UserName,
			Status:         in.Status,
			Ip:             in.Ip,
			CreateBy:       in.CreateBy,
			LastUpdateBy:   in.CreateBy,
			LastUpdateTime: time.Now(),
		})

	if err != nil {
		return nil, err
	}

	return &sys.LoginLogAddResp{}, nil
}
