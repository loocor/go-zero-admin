package logic

import (
	"context"

	"go-zero-admin/model/pms"

	"go-zero-admin/service/pms/internal/svc"
	"go-zero-admin/service/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumPicAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicAddLogic {
	return &AlbumPicAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicAddLogic) AlbumPicAdd(in *pms.AlbumPicAddReq) (*pms.AlbumPicAddResp, error) {
	_, err := l.svcCtx.PmsAlbumPicModel.Insert(
		pmsmodel.PmsAlbumPic{
			AlbumId: in.AlbumId,
			Pic:     in.Pic,
		})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumPicAddResp{}, nil
}
