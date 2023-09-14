package handler

import (
	"net/http"

	"snkrs/apps/api_gateway/internal/logic"
	"snkrs/apps/api_gateway/internal/svc"
	"snkrs/apps/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductCommentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductCommentLogic(r.Context(), svcCtx)
		resp, err := l.ProductComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
