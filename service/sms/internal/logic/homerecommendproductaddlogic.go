package logic

import (
	"context"

	"go-zero-admin/model/sms"

	"go-zero-admin/service/sms/internal/svc"
	"go-zero-admin/service/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductAddLogic {
	return &HomeRecommendProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(in *sms.HomeRecommendProductAddReq) (*sms.HomeRecommendProductAddResp, error) {
	_, err := l.svcCtx.SmsHomeRecommendProductModel.Insert(
		smsmodel.SmsHomeRecommendProduct{
			ProductId:       in.ProductId,
			ProductName:     in.ProductName,
			RecommendStatus: in.RecommendStatus,
			Sort:            in.Sort,
		})
	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendProductAddResp{}, nil
}
