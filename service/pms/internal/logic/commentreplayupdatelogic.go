package logic

import (
	"context"
	"time"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayUpdateLogic {
	return &CommentReplayUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayUpdateLogic) CommentReplayUpdate(in *pms.CommentReplayUpdateReq) (*pms.CommentReplayUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsCommentReplayModel.Update(
		pmsmodel.PmsCommentReplay{
			Id:             in.Id,
			CommentId:      in.CommentId,
			MemberNickName: in.MemberNickName,
			MemberIcon:     in.MemberIcon,
			Content:        in.Content,
			CreateTime:     CreateTime,
			Type:           in.Type,
		})
	if err != nil {
		return nil, err
	}

	return &pms.CommentReplayUpdateResp{}, nil
}
