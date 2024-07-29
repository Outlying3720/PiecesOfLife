package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	r, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return &check.CheckResp{
				Found: false,
				Price: -1,
			}, nil
		}
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: r.Price,
	}, nil
}
