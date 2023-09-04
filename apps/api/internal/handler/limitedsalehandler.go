package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"snkrs/apps/api/internal/logic"
	"snkrs/apps/api/internal/svc"
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
