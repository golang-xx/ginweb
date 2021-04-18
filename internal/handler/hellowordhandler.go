package handler

import (
	"net/http"

	"ginweb02/internal/logic"
	"ginweb02/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func helloWordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewHelloWordLogic(r.Context(), ctx)
		err := l.HelloWord()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
