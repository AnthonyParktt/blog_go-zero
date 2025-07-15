package result

type (
	baseResp struct {
		Code uint32 `json:"code"`
		Msg  string `json:"msg"`
	}

	SuccessResp struct {
		baseResp
		Data any `json:"data"`
	}

	ErrResp struct {
		baseResp
	}
)

func SuccessRespAll(data any, msg string, code uint32) *SuccessResp {
	return &SuccessResp{
		baseResp: baseResp{
			Code: code,
			Msg:  msg,
		},
		Data: data,
	}
}

func SuccessRespWithData(data any) *SuccessResp {
	return SuccessRespAll(data, "success", 200)
}

func SuccessRespWithDataMsg(data any, msg string) *SuccessResp {
	return SuccessRespAll(data, msg, 200)
}

func ErrRespAll(msg string, code uint32) *ErrResp {
	return &ErrResp{
		baseResp: baseResp{
			Code: code,
			Msg:  msg,
		},
	}
}

func ErrRespWithMsg(msg string) *ErrResp {
	return ErrRespAll(msg, 500)
}
