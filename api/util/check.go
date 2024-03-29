package util

import (
	"server/core"
	"server/core/response"

	"go.uber.org/zap"
)

var CheckUtil = checkUtil{}

// checkUtil 错误校验工具类
type checkUtil struct{}

// CheckErr 校验未知错误并抛出
func (cu checkUtil) CheckErr(err error, template string, args ...interface{}) {
	prefix := ": "
	if len(args) > 0 {
		prefix = " ,"
	}
	args = append(args, err)
	if err != nil {
		core.Logger.WithOptions(zap.AddCallerSkip(1)).Errorf(template+prefix+"err=[%+v]", args...)
		panic(response.SystemError)
	}
}

// CheckApiErr 抛出API错误
func (cu checkUtil) CheckApiErr(err error, msg string) {
	if err != nil {
		panic(response.Failed.Make(msg + ": " + err.Error()))
	}
}
