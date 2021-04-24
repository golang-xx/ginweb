package handler

import (
	"fmt"
	"net/http"

	"ginweb02/internal/logic"
	"ginweb02/internal/svc"
	"ginweb02/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func edituserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1111")
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("222",err)
			httpx.Error(w, err)
			return
		}
		fmt.Println("333")
		l := logic.NewEdituserLogic(r.Context(), ctx)
		fmt.Println("444")
		resp, err := l.Edituser(req)
		if err != nil {
			fmt.Println("555",err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
