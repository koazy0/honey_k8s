package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"honey_server/internal/common"
	"honey_server/internal/service"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(Middleware())
}

var (
	logger        = common.Logs().Cat("middleware")
	insMiddleware = sMiddleware{}
)

func Middleware() *sMiddleware {
	return &insMiddleware
}

// DefaultHandlerResponse is the default implementation of HandlerResponse.

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
