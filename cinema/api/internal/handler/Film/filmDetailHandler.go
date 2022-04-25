package Film

import (
	"net/http"

	"cinema-shop/services/film/api/internal/logic/Film"
	"cinema-shop/services/film/api/internal/svc"
	"cinema-shop/services/film/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FilmDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilmDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Film.NewFilmDetailLogic(r.Context(), svcCtx)
		resp, err := l.FilmDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
