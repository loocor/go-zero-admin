package logic

import (
	"context"
	"time"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddLogic) MemberAdd(in *ums.MemberAddReq) (*ums.MemberAddResp, error) {
	createTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	birthday, _ := time.Parse("2006-01-02 15:04:05", in.Birthday)
	l.svcCtx.UmsMemberModel.Insert(
		umsmodel.UmsMember{
			MemberLevelId:         in.MemberLevelId,
			Username:              in.Username,
			Password:              in.Password,
			Nickname:              in.Nickname,
			Phone:                 in.Phone,
			Status:                in.Status,
			CreateTime:            createTime,
			Icon:                  in.Icon,
			Gender:                in.Gender,
			Birthday:              birthday,
			City:                  in.City,
			Job:                   in.Job,
			PersonalizedSignature: in.PersonalizedSignature,
			SourceType:            in.SourceType,
			Integration:           in.Integration,
			Growth:                in.Growth,
			LuckyCount:            in.LuckyCount,
			HistoryIntegration:    in.HistoryIntegration,
		})

	return &ums.MemberAddResp{}, nil
}
