package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"bookstore/user/api/internal/config"
	"bookstore/user/api/internal/middleware"
	"bookstore/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	UserCheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	um := model.NewUserModel(conn)

	return &ServiceContext{
		Config:    c,
		UserModel: um,
		UserCheck: middleware.NewUserCheckMiddleware().Handle,
	}
}
