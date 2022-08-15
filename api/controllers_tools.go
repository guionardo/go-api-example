package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getQueryVars(ctx *gin.Context, paramNames ...string) (values []string) {
	values = make([]string, len(paramNames))
	for index, param := range paramNames {
		value := ctx.Params.ByName(param)
		if value == "" {
			value = ctx.Request.URL.Query().Get(param)
		}
		values[index] = value
	}
	return
}

func unpackVars(values []string, vars ...*string) {
	for index, value := range values {
		*vars[index] = value
	}
}

func getStrParam(ctx *gin.Context, param string) (value string, err error) {
	value = ctx.Request.URL.Query().Get(param)
	if value == "" {
		err = fmt.Errorf("Missing %s", param)
		ctx.JSON(400, err.Error())
	}
	return
}

func errorResponse(ctx *gin.Context, err error, status ...int) {
	statusCode := 500
	if len(status) > 0 {
		statusCode = status[0]
	}
	ctx.JSON(statusCode, ErrorMessage{Message: err.Error()})
}
