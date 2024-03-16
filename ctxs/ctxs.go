package ctxs

import "context"

const (
	ctxKey_UserID    ctxKey = "ctxKey_UserID"
	ctxKey_APIAuthID ctxKey = "ctxKey_APIAuthID"
	ctxKey_Trace     ctxKey = "ctxKey_Trace"
	ctxKey_CallLine  ctxKey = "ctxKey_CallLine"
)

type ctxKey string

// UserID 获取ctx中的用户标识
func UserID(ctx context.Context) uint {
	value := ctx.Value(ctxKey_UserID)
	if userID, ok := value.(uint); ok {
		return userID
	}
	return 0
}

// APIAuthID 获取ctx中的API凭证标识
func APIAuthID(ctx context.Context) uint {
	value := ctx.Value(ctxKey_APIAuthID)
	if apiAuthID, ok := value.(uint); ok {
		return apiAuthID
	}
	return 0
}

// CallLine 获取log打印行
func CallLine(ctx context.Context) string {
	value := ctx.Value(ctxKey_CallLine)
	if callLine, ok := value.(string); ok {
		return callLine
	}
	return ""
}

// WithTrace ctx添加堆栈信息
func WithTrace(ctx context.Context, trace string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_Trace, trace)
}

// WithCallLine ctx添加log打印行
func WithCallLine(ctx context.Context, callLine string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_CallLine, callLine)
}
