package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/user/model"
	"bookstore/user/rpc/user/internal/svc"
	"bookstore/user/rpc/user/user"
)

type IsUserExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserExistLogic {
	return &IsUserExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsUserExistLogic) IsUserExist(in *user.UserExistReq) (*user.UserExistReply, error) {
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
		return &user.UserExistReply{
			Exists: true,
		}, nil

	case model.ErrNotFound:
		return &user.UserExistReply{
			Exists: false,
		}, nil
	default:
		return nil, err
	}
}
