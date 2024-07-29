package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"bookstore/bookstore/rpc/bookstore"
	"bookstore/borrow/api/internal/svc"
	"bookstore/borrow/api/internal/types"
	"bookstore/borrow/model"
	"bookstore/shared"
	"bookstore/user/rpc/user/user"
)

type BorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowLogic {
	return &BorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowLogic) Borrow(userId string, req *types.BorrowReq) error {
	userInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return err
	}

	if req.ReturnPlan < time.Now().Unix() {
		return errInvalidParam
	}

	reply, err := l.svcCtx.UserRpc.IsUserExist(l.ctx, &user.UserExistReq{Id: userInt})
	if err != nil {
		if err == sqlx.ErrNotFound {
			fmt.Println("err == sqlx.ErrNotFound")
			return errUserNotFound
		}
		return err
	}

	if !reply.Exists {
		return errUserNotFound
	}

	book, err := l.svcCtx.BookstoreRpc.FindBookByName(l.ctx, &bookstore.FindBookReq{Name: req.BookName})
	if err != nil {
		if shared.IsGRPCNotFound(err) {
			fmt.Println("IsGRPCNotFound")
			return errBookNotFound
		}
		return err
	}

	_, err = l.svcCtx.BorrowSystemModel.FindOneByBookNo(book.No, model.Borrowing)
	switch err {
	case nil:
		return errBookBorrowed
	case model.ErrNotFound:
		_, err := l.svcCtx.BorrowSystemModel.Insert(l.ctx, &model.BorrowSystem{
			BookNo:         book.No,
			UserId:         userInt,
			Status:         model.Borrowing,
			ReturnPlanDate: time.Unix(req.ReturnPlan, 0),
		})
		return err
	default:
		return err
	}
}
