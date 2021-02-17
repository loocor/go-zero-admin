package logic

import (
	"context"

	"go-zero-admin/service/sms/internal/svc"
	"go-zero-admin/service/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionDeleteLogic {
	return &FlashPromotionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(in *sms.FlashPromotionDeleteReq) (*sms.FlashPromotionDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionDeleteResp{}, nil
}
