package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
)

func InitialLog() {

}

var (
	// 开启打印堆栈
	enableTrace bool
	// 模块名称
	moduleName string
)

// LogInfo 打印info日志。
func LogInfo(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.WithContext(suitCtxLine(ctx)).Info(separateArgs(args...))
	}
}

// LogWarn 打印warn日志。
func LogWarn(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.WarnLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Warn(separateArgs(args...))
	}
}

// GetStack 获取项目内部的堆栈调用信息，空格间隔。
func GetStack() string {
	sb := &strings.Builder{}
	var pc [4096]uintptr
	l := runtime.Callers(2, pc[:])
	frames := runtime.CallersFrames(pc[:l])
	for {
		frame, more := frames.Next()
		line := trimCallerLine(frame.File)
		if len(line) > 0 {
			// _, _ = fmt.Fprintf(sb, "%s\n", frame.Function)
			_, _ = fmt.Fprintf(sb, "%s:%v ", line, frame.Line)
		}
		if !more {
			break
		}
	}
	return strings.TrimSpace(sb.String())
}

// 获取除去项目名路径前部分的路径
func trimCallerLine(line string) string {
	if len(moduleName) <= 0 {
		return line
	}
	if index := strings.LastIndex(line, moduleName); index >= 0 {
		return line[index+len(moduleName)+1:]
	}
	return ""
}

// 设置调用行信息
func suitCtxLine(ctx context.Context) context.Context {
	if len(ctxs.CallLine(ctx)) > 0 {
		return ctx
	}
	_, file, line, _ := runtime.Caller(2)
	if file = trimCallerLine(file); len(file) > 0 {
		return ctxs.WithCallLine(ctx, fmt.Sprintf("%s:%d", file, line))
	}
	return ctx
}

// 格式化日志参数，空格间隔。
func separateArgs(args ...any) string {
	arr := make([]string, len(args))
	for i := range args {
		arr[i] = fmt.Sprint(args[i])
	}
	return strings.Join(arr, " ")
}
