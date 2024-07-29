package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/bookstore/model"
	"bookstore/bookstore/rpc/bookstore"
	"bookstore/bookstore/rpc/internal/svc"
	"bookstore/shared"
)

type FindBookByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBookByNameLogic {
	return &FindBookByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过书籍名称查找书籍
func (l *FindBookByNameLogic) FindBookByName(in *bookstore.FindBookReq) (*bookstore.FindBookReply, error) {
	r, err := l.svcCtx.BookstoreModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, shared.NewGRPCNotFound()
		}
		return nil, shared.NewGRPCErrorFromError(err)
	}

	return &bookstore.FindBookReply{
		No:          r.Id,
		Name:        r.Name,
		Author:      r.Author,
		PublishFate: r.PublishDate.Time.Format("2006-12-25"),
	}, nil
}
