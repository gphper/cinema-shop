package handler

import (
	"cinema-shop/common/errorxx"
	"cinema-shop/services/usercenter/api/internal/logic"
	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {

			if reflect.TypeOf(err).Name() == "ApiCustomError" {
				er := err.(errorxx.ApiCustomError)
				httpx.OkJson(w, er)
			} else {
				httpx.Error(w, err)
			}

		} else {
			httpx.OkJson(w, resp)
		}
	}
}
