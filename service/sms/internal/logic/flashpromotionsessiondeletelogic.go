package logic

import (
	"context"

	"go-zero-admin/service/sms/internal/svc"
	"go-zero-admin/service/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionDeleteLogic {
	return &FlashPromotionSessionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionDeleteLogic) FlashPromotionSessionDelete(in *sms.FlashPromotionSessionDeleteReq) (*sms.FlashPromotionSessionDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionSessionModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionSessionDeleteResp{}, nil
}
