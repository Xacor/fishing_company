package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomHttpFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%-12v %-6v %-30v %-3v %-10v %v\n",
		param.ClientIP,
		param.Method,
		param.Path,
		param.StatusCode,
		param.Latency,
		param.ErrorMessage,
	)
}
