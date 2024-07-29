package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"bookstore/bookstore/model"
	"bookstore/bookstore/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	BookstoreModel model.BookstoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		BookstoreModel: model.NewBookstoreModel(conn),
	}
}
