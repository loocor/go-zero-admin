package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthLogic {
	return AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req types.TokenReq) (*types.TokenResp, error) {
	// TODO: check user id & password
	payloads := map[string]interface{}{
		"UserId":   req.UserId,
		"Password": req.Password,
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := l.GenToken(now, l.svcCtx.Config.Auth.AccessSecret, payloads, accessExpire)

	if err != nil {
		return nil, err
	}

	return &types.TokenResp{
		CommonResp: types.CommonResp{
			Code:    0,
			Message: "done",
		},
		Token: types.JwtToken{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

func (l *AuthLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
