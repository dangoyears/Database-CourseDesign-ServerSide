package dbcd

import (
	"log"
	"runtime"
)

// Trace 跟踪操作中的错误信息。
// 向日志文件中输出两行，第一行是当前正在执行的函数，第二行是错误信息。
func Trace(v ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	caller := frame.Function

	log.Print(caller, " ")
	log.Print(v...)
}
