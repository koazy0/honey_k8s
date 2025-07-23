// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// AccessKeyAuth 通过accesskey进行认证,后面会进行修改
		AccessKeyAuth(r *ghttp.Request)
		// JWTAuth 使用 JWT 进行认证
		JWTAuth(r *ghttp.Request)
		// MiddleLogger 用于记录谁进行了操作，放在中间件当中
		MiddleLogger(r *ghttp.Request)
		// CORS allows Cross-origin resource sharing.
		CORS(r *ghttp.Request)
		// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
		MiddlewareHandlerResponse(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
