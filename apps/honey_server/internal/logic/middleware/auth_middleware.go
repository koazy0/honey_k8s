package middleware

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/samber/lo"

	"honey_server/internal/service"
	"net/http"
	"strings"
)

// AccessKeyAuth 通过accesskey进行认证,后面会进行修改
func (s *sMiddleware) AccessKeyAuth(r *ghttp.Request) {
	// 从请求头获取 AK 值
	ak := r.GetHeader("Access-Key") // 或 r.Header.Get("AK")
	configAK, _ := g.Cfg().Get(context.Background(), "server.AccessKey")
	whiteList, _ := g.Cfg().Get(context.Background(), "white_list")

	_, exist := lo.Find(whiteList.Array(), func(item interface{}) bool {
		return r.Router.Uri == item.(string)
	})

	// 如果路由不在白名单，或者AK为空或不匹配
	if !exist && ak != configAK.String() {
		r.Response.WriteStatusExit(
			http.StatusUnauthorized,
			ghttp.DefaultHandlerResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid access key",
			},
		)
		return
	}
	// 验证通过，继续后续处理
	r.Middleware.Next()
}

// JWTAuth 使用 JWT 进行认证
func (s *sMiddleware) JWTAuth(r *ghttp.Request) {
	authHeader := r.GetHeader("Authorization") // 或 r.Header.Get("AK")
	whiteList, _ := g.Cfg().Get(context.Background(), "white_list")

	// 如果在白名单直接进入下一个中间件
	_, exist := lo.Find(whiteList.Array(), func(item interface{}) bool {
		return r.Router.Uri == item.(string)
	})
	if exist {
		r.Middleware.Next()
		return
	}

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		//不用下面默认的回复方式
		//r.Response.WriteStatusExit(
		//	http.StatusUnauthorized,
		//	ghttp.DefaultHandlerResponse{
		//		Code:    http.StatusUnauthorized,
		//		Message: "Missing or invalid Authorization header",
		//	},
		//)
		err := gerror.NewCode(gcode.CodeNotAuthorized, "缺少或非法的 Authorization 头")
		r.SetError(err)
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := service.Jwt().ParseToken(tokenString)
	if err != nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, err.Error())
		r.SetError(err)
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	// 认证通过，设置用户信息上下文
	r.SetCtxVar("user_id", userID)
	r.Middleware.Next()
}
