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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
	"gitee.com/CertificateAndSigningManageSystem/common/errs"
)

// InitialLog 初始化日志
func InitialLog(logDir, module string, maxAge, rotationTime time.Duration, debug bool) {
	err := os.MkdirAll(logDir, 0777)
	if err != nil {
		panic("make log dir error " + err.Error())
	}
	baseLogFile := path.Join(logDir, "csms")
	debugWriter, err := rotatelogs.New(
		baseLogFile+"_debug_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic("rotate logs error " + err.Error())
	}
	infoWriter, err := rotatelogs.New(
		baseLogFile+"_info_%Y%m%d%H%M.log",
		// rotatelogs.WithLinkName(baseLogFile),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic("rotate logs error " + err.Error())
	}
	errorWriter, err := rotatelogs.New(
		baseLogFile+"_error_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic("rotate logs error " + err.Error())
	}
	formatter := &logFormatter{}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debugWriter, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  io.MultiWriter(infoWriter, debugWriter),
		logrus.WarnLevel:  io.MultiWriter(infoWriter, debugWriter),
		logrus.ErrorLevel: io.MultiWriter(errorWriter, infoWriter, debugWriter),
		logrus.FatalLevel: io.MultiWriter(errorWriter, infoWriter, debugWriter),
	}, formatter)
	logrus.SetFormatter(formatter)
	logrus.AddHook(lfHook)
	if debug {
		enableTrace = true
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	moduleName = module
}

var (
	// 开启打印堆栈
	enableTrace bool
	// 模块名称
	moduleName string
)

type GormLogFormatter struct{}

type TusClientLogger struct{}

type logFormatter struct {
	TimestampFormat string
	LogFormat       string
}

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

func (f *GormLogFormatter) Printf(s string, args ...any) {
	if len(args) < 4 {
		Warnf(ctxs.NewCtx("Printf"), "gorm log unknown fmt "+s, args)
		return
	}
	ctx := ctxs.WithCallLine(nil, trimCallerLine(fmt.Sprint(args[0])))
	err, _ := args[1].(error)
	if err != nil {
		Errorf(ctx, "%s [%.3fms] %s", err, args[2], args[4])
	} else {
		Infof(ctx, "[%.3fms] [rows:%v] %s", args[1], args[2], args[3])
	}
}

func (l *TusClientLogger) Error(ctx context.Context, msg string) {
	Error(ctx, msg)
}

func (l *TusClientLogger) Info(ctx context.Context, msg string) {
	Info(ctx, msg)
}

func (l *TusClientLogger) Warn(ctx context.Context, msg string) {
	Warn(ctx, msg)
}

func (l *TusClientLogger) Debug(ctx context.Context, msg string) {
	Debug(ctx, msg)
}

// Format 格式化日志
func (formatter *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	sb := strings.Builder{}
	sb.WriteString(entry.Time.Format("2006-01-02 15:04:05.000 "))
	sb.WriteString(strings.ToUpper(entry.Level.String()))
	sb.WriteString(" ")
	sb.WriteString(moduleName)
	sb.WriteString(" ")
	ctx := entry.Context
	if rid := ctxs.RequestId(ctx); len(rid) > 0 {
		sb.WriteString(rid)
		sb.WriteString(" ")
	}
	if line := ctxs.CallLine(ctx); len(line) > 0 {
		sb.WriteString(line)
		sb.WriteString(" ")
	} else if entry.HasCaller() {
		if file := trimCallerLine(entry.Caller.File); len(file) > 0 {
			sb.WriteString(fmt.Sprintf("%s:%d ", file, entry.Caller.Line))
		}
	}
	if fn := ctxs.Func(ctx); len(fn) > 0 {
		sb.WriteString(fn)
		sb.WriteString(" ")
	}
	if userId := ctxs.UserId(ctx); userId > 0 {
		sb.WriteString(strconv.Itoa(int(userId)))
		sb.WriteString(" ")
	}
	if auth := ctxs.APIAuthId(ctx); auth > 0 {
		sb.WriteString(strconv.Itoa(int(auth)))
		sb.WriteString(" ")
	}
	if ip := ctxs.RequestIP(ctx); len(ip) > 0 {
		sb.WriteString(ip)
		sb.WriteString(" ")
	}
	if ph := ctxs.RequestPath(ctx); len(ph) > 0 {
		sb.WriteString(ph)
		sb.WriteString(" ")
	}
	if trace := ctxs.Trace(ctx); len(trace) > 0 {
		sb.WriteString("[")
		sb.WriteString(trace)
		sb.WriteString("] ")
	}
	sb.WriteString(fmt.Sprintf("-- %s", entry.Message))
	msg := strings.ReplaceAll(sb.String(), "\r\n", `\r\n`)
	msg = strings.ReplaceAll(msg, "\n", `\n`)

	return []byte(msg + "\n"), nil
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
		switch v := args[i].(type) {
		case json.RawMessage:
			arr[i] = string(v)
		case string:
			arr[i] = v
		case *errs.Error:
			arr[i] = v.WrappedErr.Error()
		case error:
			arr[i] = v.Error()
		case []byte:
			arr[i] = string(v)
		case []rune:
			arr[i] = string(v)
		default:
			bs, err := json.Marshal(v)
			if err != nil {
				arr[i] = fmt.Sprint(v)
			} else {
				arr[i] = string(bs)
			}
		}
	}
	return strings.Join(arr, " ")
}
