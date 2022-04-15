package handler

import (
	"cinema-shop/common/errorxx"
	"cinema-shop/services/usercenter/api/internal/logic"
	"cinema-shop/services/usercenter/api/internal/svc"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
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
