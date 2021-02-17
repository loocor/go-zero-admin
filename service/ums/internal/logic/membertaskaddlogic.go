package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskAddLogic {
	return &MemberTaskAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(in *ums.MemberTaskAddReq) (*ums.MemberTaskAddResp, error) {
	_, err := l.svcCtx.UmsMemberTaskModel.Insert(
		umsmodel.UmsMemberTask{
			Name:        in.Name,
			Growth:      in.Growth,
			Integration: in.Integration,
			Type:        in.Type,
		})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskAddResp{}, nil
}
