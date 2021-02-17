package logic

import (
	"context"
	"time"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryAddLogic {
	return &GrowthChangeHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(in *ums.GrowthChangeHistoryAddReq) (*ums.GrowthChangeHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsGrowthChangeHistoryModel.Insert(
		umsmodel.UmsGrowthChangeHistory{
			MemberId:    in.MemberId,
			CreateTime:  CreateTime,
			ChangeType:  in.ChangeType,
			ChangeCount: in.ChangeCount,
			OperateMan:  in.OperateMan,
			OperateNote: in.OperateNote,
			SourceType:  in.SourceType,
		})
	if err != nil {
		return nil, err
	}

	return &ums.GrowthChangeHistoryAddResp{}, nil
}
