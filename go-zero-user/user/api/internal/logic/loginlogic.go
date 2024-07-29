package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/user/api/internal/svc"
	"bookstore/user/api/internal/types"
	"bookstore/user/model"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, second, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + second
	claims["at"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.UserReply, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	switch err {
	case nil:
		if userInfo.Password != req.Password {
			return nil, errorIncorrectPassword
		}

		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
		if err != nil {
			return nil, err
		}

		logx.Info(jwtToken)

		return &types.UserReply{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
			JwtToken: types.JwtToken{
				AccessToken:  jwtToken,
				AccessExpire: now + accessExpire,
				RefreshAfter: now + accessExpire/2,
			},
		}, nil

	case model.ErrNotFound:
		return nil, errorUsernameUnRegister

	default:
		return nil, err
	}
}
