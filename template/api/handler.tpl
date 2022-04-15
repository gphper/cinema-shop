package {{.PkgName}}

import (
	"net/http"
	"reflect"
	"cinema-shop/common/errorxx"
	{{if .After1_1_10}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {

			if reflect.TypeOf(err).Name() == "ApiCustomError" {
				er := err.(errorxx.ApiCustomError)
				httpx.OkJson(w, er)
			} else {
				httpx.Error(w, err)
			}

		} else {
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
