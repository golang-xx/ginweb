package handler

import (
	"net/http"

	"ginweb02/internal/logic"
	"ginweb02/internal/svc"
	"ginweb02/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func verifyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVerifyLogic(r.Context(), ctx)
		resp, err := l.Verify(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
