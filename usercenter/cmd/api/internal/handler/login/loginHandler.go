package login

import (
	"go-zero_less/usercenter/cmd/api/internal/logic/login"
	"go-zero_less/usercenter/cmd/api/internal/svc"
	"go-zero_less/usercenter/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			//result.HttpResult(r, w, resp, err)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//result.HttpResult(r, w, resp, err)
			//httpx.OkJsonCtx(r.Context(), w, resp)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
