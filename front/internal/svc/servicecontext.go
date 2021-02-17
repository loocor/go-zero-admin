package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/front/internal/config"
	"go-zero-admin/service/oms/omsclient"
	"go-zero-admin/service/pms/pmsclient"
	"go-zero-admin/service/sms/smsclient"
	"go-zero-admin/service/sys/sysclient"
	"go-zero-admin/service/ums/umsclient"
)

type ServiceContext struct {
	Config config.Config

	Sys sysclient.Sys
	Ums umsclient.Ums
	Pms pmsclient.Pms
	Oms omsclient.Oms
	Sms smsclient.Sms
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Sys: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums: umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms: pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms: omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms: smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
	}
}
