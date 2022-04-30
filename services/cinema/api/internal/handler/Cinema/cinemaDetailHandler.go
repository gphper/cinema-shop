package Cinema

import (
	"net/http"

	"cinema-shop/services/cinema/api/internal/logic/Cinema"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CinemaDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CinemaDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Cinema.NewCinemaDetailLogic(r.Context(), svcCtx)
		resp, err := l.CinemaDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
