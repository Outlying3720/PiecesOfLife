package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

const (
	_ = iota
	Borrowing
	Return
)

var _ BorrowSystemModel = (*customBorrowSystemModel)(nil)

type (
	// BorrowSystemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBorrowSystemModel.
	BorrowSystemModel interface {
		borrowSystemModel
		withSession(session sqlx.Session) BorrowSystemModel
		FindOneByUserAndBookNo(userId int64, bookNo string) (*BorrowSystem, error)
		FindOneByBookNo(bookNo string, status int) (*BorrowSystem, error)
	}

	customBorrowSystemModel struct {
		*defaultBorrowSystemModel
	}
)

// NewBorrowSystemModel returns a model for the database table.
func NewBorrowSystemModel(conn sqlx.SqlConn) BorrowSystemModel {
	return &customBorrowSystemModel{
		defaultBorrowSystemModel: newBorrowSystemModel(conn),
	}
}

func (m *customBorrowSystemModel) withSession(session sqlx.Session) BorrowSystemModel {
	return NewBorrowSystemModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultBorrowSystemModel) FindOneByUserAndBookNo(userId int64, bookNo string) (*BorrowSystem, error) {
	query := "select " + borrowSystemRows + " from " + m.table + " where user_id=? and book_no = ? limit 1"
	var resp BorrowSystem
	err := m.conn.QueryRow(&resp, query, userId, bookNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBorrowSystemModel) FindOneByBookNo(bookNo string, status int) (*BorrowSystem, error) {
	query := "select " + borrowSystemRows + " from " + m.table + " where user_id=? and status = ? limit 1"
	var resp *BorrowSystem
	err := m.conn.QueryRow(resp, query, bookNo, status)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
