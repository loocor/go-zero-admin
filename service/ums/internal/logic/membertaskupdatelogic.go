package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskUpdateLogic {
	return &MemberTaskUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(in *ums.MemberTaskUpdateReq) (*ums.MemberTaskUpdateResp, error) {
	err := l.svcCtx.UmsMemberTaskModel.Update(
		umsmodel.UmsMemberTask{
			Id:          in.Id,
			Name:        in.Name,
			Growth:      in.Growth,
			Integration: in.Integration,
			Type:        in.Type,
		})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskUpdateResp{}, nil
}
