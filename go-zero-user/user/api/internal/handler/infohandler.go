package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"bookstore/user/api/internal/logic"
	"bookstore/user/api/internal/svc"
)

func infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")
		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(userId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
