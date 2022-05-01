package Screen

import (
	"net/http"

	"cinema-shop/services/cinema/api/internal/logic/Screen"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ScreenDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScreenDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Screen.NewScreenDetailLogic(r.Context(), svcCtx)
		resp, err := l.ScreenDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
