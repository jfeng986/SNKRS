package handler

import (
	"net/http"

	"snkrs/apps/api_gateway/internal/logic"
	"snkrs/apps/api_gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LimitedSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLimitedSaleLogic(r.Context(), svcCtx)
		resp, err := l.LimitedSale()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
