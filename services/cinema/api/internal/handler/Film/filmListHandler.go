package Film

import (
	"net/http"

	"cinema-shop/services/cinema/api/internal/logic/Film"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FilmListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilmListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Film.NewFilmListLogic(r.Context(), svcCtx)
		resp, err := l.FilmList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
