package svc

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"bookstore/bookstore/rpc/bookstoreclient"
	"bookstore/borrow/api/internal/config"
	"bookstore/borrow/model"
	"bookstore/user/rpc/user/userclient"
)

type ServiceContext struct {
	Config            config.Config
	BorrowSystemModel model.BorrowSystemModel
	UserRpc           userclient.User
	BookstoreRpc      bookstoreclient.Bookstore
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ur := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	br := bookstoreclient.NewBookstore(zrpc.MustNewClient(c.BookstoreRpc))

	fmt.Println(&br)
	return &ServiceContext{
		Config:            c,
		BorrowSystemModel: model.NewBorrowSystemModel(conn),
		UserRpc:           ur,
		BookstoreRpc:      br,
	}
}
