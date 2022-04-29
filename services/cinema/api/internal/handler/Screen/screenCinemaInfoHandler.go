package Screen

import (
	"net/http"

	"cinema-shop/services/cinema/api/internal/logic/Screen"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ScreenCinemaInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScreenCinemaInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			httpx.Error(w, err)
			return
		}

		l := Screen.NewScreenCinemaInfoLogic(r.Context(), svcCtx)
		resp, err := l.ScreenCinemaInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
