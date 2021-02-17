package logic

import (
	"context"

	"go-zero-admin/model/sms"

	"go-zero-admin/service/sms/internal/svc"
	"go-zero-admin/service/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendSubjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectUpdateLogic {
	return &HomeRecommendSubjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(in *sms.HomeRecommendSubjectUpdateReq) (*sms.HomeRecommendSubjectUpdateResp, error) {
	err := l.svcCtx.SmsHomeRecommendSubjectModel.Update(
		smsmodel.SmsHomeRecommendSubject{
			Id:              in.Id,
			SubjectId:       in.SubjectId,
			SubjectName:     in.SubjectName,
			RecommendStatus: in.RecommendStatus,
			Sort:            in.Sort,
		})
	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendSubjectUpdateResp{}, nil
}