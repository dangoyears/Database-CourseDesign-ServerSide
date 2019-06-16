package dbcd

import (
	"log"
	"runtime"
)

// Trace 跟踪操作中的错误信息。
func (engine *Engine) Trace(v ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	caller := frame.Function

	log.Print(caller, " ")
	log.Print(v...)
}
