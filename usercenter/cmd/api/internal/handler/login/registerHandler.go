package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-zero_less/usercenter/cmd/api/internal/logic/login"
	"go-zero_less/usercenter/cmd/api/internal/svc"
	"go-zero_less/usercenter/cmd/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegiesterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
