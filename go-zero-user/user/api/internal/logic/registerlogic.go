package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/user/api/internal/svc"
	"bookstore/user/api/internal/types"
	"bookstore/user/model"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	_, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	if err == nil {
		return errorDuplicateUsername
	}

	_, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err == nil {
		return errorDuplicateMobile
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:     req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   "男",
		Nickname: "anonymous",
	})

	return err
}
