package logic

import (
	"context"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressDeleteLogic {
	return MemberAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressDeleteLogic) MemberAddressDelete(req types.MemberAddressDeleteReq) (*types.MemberAddressDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberAddressDeleteResp{}, nil
}
