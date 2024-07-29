package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BookstoreModel = (*customBookstoreModel)(nil)

type (
	// BookstoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookstoreModel.
	BookstoreModel interface {
		bookstoreModel
		withSession(session sqlx.Session) BookstoreModel
	}

	customBookstoreModel struct {
		*defaultBookstoreModel
	}
)

// NewBookstoreModel returns a model for the database table.
func NewBookstoreModel(conn sqlx.SqlConn) BookstoreModel {
	return &customBookstoreModel{
		defaultBookstoreModel: newBookstoreModel(conn),
	}
}

func (m *customBookstoreModel) withSession(session sqlx.Session) BookstoreModel {
	return NewBookstoreModel(sqlx.NewSqlConnFromSession(session))
}
