package logic

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/user/api/internal/svc"
	"bookstore/user/api/internal/types"
	"bookstore/user/model"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(userId string) (resp *types.UserReply, err error) {
	userInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userInt)
	switch err {
	case nil:
		return &types.UserReply{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
			Nickname: userInfo.Nickname,
			Gender:   "ok",
		}, nil

	case model.ErrNotFound:
		return nil, errorUserNotFound
	default:
		return nil, err
	}
}
