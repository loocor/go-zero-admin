package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagAddLogic {
	return &MemberTagAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagAddLogic) MemberTagAdd(in *ums.MemberTagAddReq) (*ums.MemberTagAddResp, error) {
	_, err := l.svcCtx.UmsMemberTagModel.Insert(
		umsmodel.UmsMemberTag{
			Name:              in.Name,
			FinishOrderCount:  in.FinishOrderCount,
			FinishOrderAmount: float64(in.FinishOrderAmount),
		})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTagAddResp{}, nil
}
