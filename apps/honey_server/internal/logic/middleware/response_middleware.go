package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"net/http"
	"strconv"
)

type DefaultHandlerResponse struct {
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Status    int         `json:"status"`
	Code      string      `json:"code"`
	Timestamp int64       `json:"timestamp"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	if msg == "" {
		msg = "success"
	}
	r.Response.WriteJson(DefaultHandlerResponse{
		Code:      strconv.Itoa(code.Code()),
		Status:    r.Response.Status,
		Message:   msg,
		Result:    res,
		Timestamp: gtime.Now().UnixMilli(),
	})
}
