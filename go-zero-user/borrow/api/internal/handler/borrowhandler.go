package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"bookstore/borrow/api/internal/logic"

	"bookstore/borrow/api/internal/svc"
	"bookstore/borrow/api/internal/types"
)

func borrowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBorrowLogic(r.Context(), svcCtx)
		err := l.Borrow(r.Header.Get("x-user-id"), &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
