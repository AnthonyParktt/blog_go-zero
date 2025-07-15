package result

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HttpResult(r *http.Request, w http.ResponseWriter, data interface{}, err error) {
	if err == nil {
		r := SuccessRespWithData(data)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		resp := ErrRespWithMsg(err.Error())
		httpx.WriteJson(w, http.StatusOK, resp)
	}
}
