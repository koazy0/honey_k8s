package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddleLogger 用于记录谁进行了操作，放在中间件当中
func (s *sMiddleware) MiddleLogger(r *ghttp.Request) {
	r.Middleware.Next()

	//进行操作日志处理
	//todo 后续考虑要不要把这段进行入库处理，暂时先记录成日志
	result := r.GetError().Error()
	if result == "" {
		result = "success"
	}
	logger.Infof("userID:%s, active: %s, result:%s ", r.GetCtxVar("user_id").String(), r.Router.Uri, result)

}
