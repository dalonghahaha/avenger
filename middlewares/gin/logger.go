package gin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"avenger/components/logger"
)

func Logger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	ctx.Next()
	end := time.Now()
	latency := end.Sub(start)
	if len(ctx.Errors) > 0 {
		for _, e := range ctx.Errors.Errors() {
			logger.Error(e)
		}
	}
	content := fmt.Sprintf("%d %s %s %s %s %dms",
		ctx.Writer.Status(),
		ctx.Request.Method,
		path,
		query,
		ctx.ClientIP(),
		latency.Microseconds(),
	)
	logger.Debug(content)
}
