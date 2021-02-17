package logic

import (
	"context"

	"go-zero-admin/model/ums"

	"go-zero-admin/service/ums/internal/svc"
	"go-zero-admin/service/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelUpdateLogic {
	return &MemberLevelUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelUpdateLogic) MemberLevelUpdate(in *ums.MemberLevelUpdateReq) (*ums.MemberLevelUpdateResp, error) {
	err := l.svcCtx.UmsMemberLevelModel.Update(
		umsmodel.UmsMemberLevel{
			Id:                    in.Id,
			Name:                  in.Name,
			GrowthPoint:           in.GrowthPoint,
			DefaultStatus:         in.DefaultStatus,
			FreeFreightPoint:      float64(in.FreeFreightPoint),
			CommentGrowthPoint:    in.CommentGrowthPoint,
			PriviledgeFreeFreight: in.PriviledgeFreeFreight,
			PriviledgeSignIn:      in.PriviledgeSignIn,
			PriviledgeComment:     in.PriviledgeComment,
			PriviledgePromotion:   in.PriviledgePromotion,
			PriviledgeMemberPrice: in.PriviledgeMemberPrice,
			PriviledgeBirthday:    in.PriviledgeBirthday,
			Note:                  in.Note,
		})
	if err != nil {
		return nil, err
	}

	return &ums.MemberLevelUpdateResp{}, nil
}
