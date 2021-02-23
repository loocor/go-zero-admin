package logic

import (
	"context"

	"zdmin/service/sys/internal/svc"
	"zdmin/service/sys/sys"

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
	// todo: add your logic here and delete this line

	return &sys.LoginLogAddResp{}, nil
}
