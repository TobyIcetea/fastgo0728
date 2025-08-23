package middleware

import (
	"github.com/TobyIcetea/fastgo/internal/pkg/contextx"
	"github.com/TobyIcetea/fastgo/internal/pkg/core"
	"github.com/TobyIcetea/fastgo/internal/pkg/errorsx"
	"github.com/TobyIcetea/fastgo/pkg/token"
	"github.com/gin-gonic/gin"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法
// 如果合法，浙江 token 中的 sub 作为 <用户名> 存放在 gin.Context 的 XUsernameKey 键中
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		userID, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errorsx.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		// 将用户 ID 和用户名注入到上下文中
		ctx := contextx.WithUserID(c.Request.Context(), userID)
		c.Request = c.Request.WithContext(ctx)

		// 继续后续的操作
		c.Next()
	}
}
