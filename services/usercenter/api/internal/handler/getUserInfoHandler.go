package handler

import (
	"net/http"

	"cinema-shop/services/usercenter/api/internal/logic"
	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func get_user_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGet_user_infoLogic(r.Context(), svcCtx)
		resp, err := l.Get_user_info(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
