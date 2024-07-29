package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"bookstore/bookstore/rpc/bookstore"
	"bookstore/borrow/api/internal/svc"
	"bookstore/borrow/api/internal/types"
	"bookstore/borrow/model"
)

type ReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnLogic {
	return &ReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnLogic) Return(userId string, req *types.ReturnReq) error {
	userInt, err := strconv.ParseInt(fmt.Sprintf("%v", userId), 10, 64)
	if err != nil {
		return err
	}

	book, err := l.svcCtx.BookstoreRpc.FindBookByName(l.ctx, &bookstore.FindBookReq{Name: req.BookName})
	if err != nil {
		return err
	}

	fmt.Println(book)

	info, err := l.svcCtx.BorrowSystemModel.FindOneByUserAndBookNo(userInt, book.No)

	fmt.Println(info, err)
	switch err {
	case nil:
		if info.Status == model.Return {
			return errBookReturn
		}
		info.ReturnDate = time.Now().Unix()
		info.Status = model.Return
		err = l.svcCtx.BorrowSystemModel.Update(l.ctx, info)
		return err
	case model.ErrNotFound:
		return errUserReturn
	default:
		return err
	}
}
