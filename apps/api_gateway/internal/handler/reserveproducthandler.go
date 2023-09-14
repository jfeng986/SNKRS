package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"snkrs/apps/api_gateway/internal/logic"
	"snkrs/apps/api_gateway/internal/svc"
	"snkrs/apps/api_gateway/internal/types"
)

func ReserveProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReserveProductRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewReserveProductLogic(r.Context(), svcCtx)
		resp, err := l.ReserveProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
