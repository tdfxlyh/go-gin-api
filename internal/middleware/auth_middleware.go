package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"net/http"
	"strings"
)

// AuthMiddleware 权限中间件: 用户登录才可以通过
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, output.NewErrorResp(output.StatusCodeNotLoggedIn, ""))
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:] //截取字符

		token, claims, err := caller.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, output.NewErrorResp(output.StatusCodeNotLoggedIn, ""))
			ctx.Abort()
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserID

		var user models.User
		caller.LyhTestDB.Debug().Table("user").Where("uid=?", userId).First(&user)

		// 验证用户是否存在
		if user.UID == 0 {
			ctx.JSON(http.StatusUnauthorized, output.NewErrorResp(output.StatusCodeNotLoggedIn, "用户不存在"))
			ctx.Abort()
			return
		}

		// 用户是否被注销
		if user.Status != 0 {
			ctx.JSON(http.StatusUnauthorized, output.NewErrorResp(output.StatusCodeNotLoggedIn, "用户已注销"))
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
