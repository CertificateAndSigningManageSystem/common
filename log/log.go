/*
 * Copyright (c) 2023 ivfzhou
 * common is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

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

type GormLogFormatter struct{}

// Infof 打印info日志。
func Infof(ctx context.Context, f string, args ...any) {
	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.WithContext(suitCtxLine(ctx)).Infof(f, args...)
	}
}

// Info 打印info日志。
func Info(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.WithContext(suitCtxLine(ctx)).Info(separateArgs(args...))
	}
}

// Warn 打印warn日志。
func Warn(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.WarnLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Warn(separateArgs(args...))
	}
}

// Warnf 打印warn日志。
func Warnf(ctx context.Context, f string, args ...any) {
	if logrus.IsLevelEnabled(logrus.WarnLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Warnf(f, args...)
	}
}

// Error 打印error日志。
func Error(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.ErrorLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Error(separateArgs(args...))
	}
}

// Errorf 打印error日志。
func Errorf(ctx context.Context, f string, args ...any) {
	if logrus.IsLevelEnabled(logrus.ErrorLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Errorf(f, args...)
	}
}

// Debug 打印debug日志。
func Debug(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.WithContext(suitCtxLine(ctx)).Debug(separateArgs(args...))
	}
}

// Debugf 打印debug日志。
func Debugf(ctx context.Context, f string, args ...any) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.WithContext(suitCtxLine(ctx)).Debugf(f, args...)
	}
}

// Fatal 打印fatal日志，之后程序退出。
func Fatal(ctx context.Context, args ...any) {
	if logrus.IsLevelEnabled(logrus.FatalLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Fatal(separateArgs(args...))
	}
}

// Fatalf 打印fatal日志，之后程序退出。
func Fatalf(ctx context.Context, f string, args ...any) {
	if logrus.IsLevelEnabled(logrus.FatalLevel) {
		if ctx = suitCtxLine(ctx); enableTrace {
			ctx = ctxs.WithTrace(ctx, GetStack())
		}
		logrus.WithContext(ctx).Fatalf(f, args...)
	}
}

// ErrorIf 如果err不为nil，则log
func ErrorIf(ctx context.Context, err error) {
	if err != nil {
		if logrus.IsLevelEnabled(logrus.ErrorLevel) {
			if ctx = suitCtxLine(ctx); enableTrace {
				ctx = ctxs.WithTrace(ctx, GetStack())
			}
			logrus.WithContext(ctx).Error(err.Error())
		}
	}
}

// FatalIfError 如果err不为nil，则log
func FatalIfError(ctx context.Context, err error) {
	if err != nil {
		if logrus.IsLevelEnabled(logrus.ErrorLevel) {
			if ctx = suitCtxLine(ctx); enableTrace {
				ctx = ctxs.WithTrace(ctx, GetStack())
			}
			logrus.WithContext(ctx).Fatal(err.Error())
		}
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

func (f *GormLogFormatter) Printf(_ string, args ...any) {
	ctx := ctxs.WithCallLine(nil, trimCallerLine(fmt.Sprint(args[0])))
	err, _ := args[1].(error)
	if err != nil {
		Errorf(ctx, "%s [%.3fms] %s", err, args[2], args[len(args)-1])
	} else {
		Infof(ctx, "[%.3fms] [rows:%v] %s", args[1], args[len(args)-2], args[len(args)-1])
	}
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