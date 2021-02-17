package logic

import (
	"context"

	"go-zero-admin/model/oms"

	"go-zero-admin/service/oms/internal/svc"
	"go-zero-admin/service/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompanyAddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressUpdateLogic {
	return &CompanyAddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressUpdateLogic) CompanyAddressUpdate(in *oms.CompanyAddressUpdateReq) (*oms.CompanyAddressUpdateResp, error) {
	err := l.svcCtx.OmsCompanyAddressModel.Update(
		omsmodel.OmsCompanyAddress{
			Id:            in.Id,
			AddressName:   in.AddressName,
			SendStatus:    in.SendStatus,
			ReceiveStatus: in.ReceiveStatus,
			Name:          in.Name,
			Phone:         in.Phone,
			Province:      in.Province,
			City:          in.City,
			Region:        in.Region,
			DetailAddress: in.DetailAddress,
		})

	if err != nil {
		return nil, err
	}

	return &oms.CompanyAddressUpdateResp{}, nil
}