package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = 0
		body.Msg = err.Error()
	} else {
		body.Code = 200
		body.Msg = "success"
	}
	httpx.OkJson(w, body)
}
