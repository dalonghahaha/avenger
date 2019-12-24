package gin

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"avenger/components/logger"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic:", err)
			logger.Error("stack:", string(debug.Stack()))
			ctx.String(http.StatusBadGateway, "服务器异常")
		}
	}()
	ctx.Next()
}
