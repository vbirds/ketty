package ketty

import (
	"github.com/yyzybb537/ketty/log"
)

func SetLog(l log.LogI) {
	log.SetLog(l)
}

func GetLog(opt ... interface{}) log.LogI {
	return log.GetLog(opt...)
}

var Indent log.LogFormatOptions = log.Indent

func LogFormat(req interface{}, opt ... log.LogFormatOptions) string {
	return log.LogFormat(req, opt...)
}

