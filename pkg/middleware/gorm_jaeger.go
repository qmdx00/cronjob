package middleware
//
//import (
//	"context"
//	"fmt"
//	"github.com/opentracing/opentracing-go"
//	"github.com/opentracing/opentracing-go/ext"
//	"gorm.io/gorm"
//)
//
//const (
//	jaegerContextKey = "jeager:context"
//)
//
//func RegisterGormJaeger(db *gorm.DB) {
//	spanTypePrefix := fmt.Sprintf("gorm.%s.", driverName)
//	querySpanType := spanTypePrefix + "query"
//	execSpanType := spanTypePrefix + "exec"
//
//	type params struct {
//		spanType  string
//		processor func() *gorm.CallbackProcessor
//	}
//	callbacks := map[string]params{
//		"gorm:create": {
//			spanType:  execSpanType,
//			processor: func() *gorm.CallbackProcessor { return db.Callback().Create() },
//		},
//		"gorm:delete": {
//			spanType:  execSpanType,
//			processor: func() *gorm.CallbackProcessor { return db.Callback().Delete() },
//		},
//		"gorm:query": {
//			spanType:  querySpanType,
//			processor: func() *gorm.CallbackProcessor { return db.Callback().Query() },
//		},
//		"gorm:update": {
//			spanType:  execSpanType,
//			processor: func() *gorm.CallbackProcessor { return db.Callback().Update() },
//		},
//		"gorm:row_query": {
//			spanType:  querySpanType,
//			processor: func() *gorm.CallbackProcessor { return db.Callback().RowQuery() },
//		},
//	}
//	for name, param := range callbacks {
//		param.processor().Before(name).Register(
//			fmt.Sprintf("%s:before:%s", "jaeger", name),
//			newBeforeCallback(param.spanType),
//		)
//		param.processor().After(name).Register(
//			fmt.Sprintf("%s:after:%s", "jaeger", name),
//			newAfterCallback(),
//		)
//	}
//}
//
//func newBeforeCallback(spanType string) func(*gorm.Scope) {
//	return func(scope *gorm.Scope) {
//		ctx, ok := scopeContext(scope)
//		if !ok {
//			return
//		}
//		//新增链路追踪
//		span, ctx := opentracing.StartSpanFromContext(ctx, spanType, opentracing.Tag{Key: string(ext.Component), Value: "Gorm operation"})
//		if span.Tracer() == nil {
//			span.Finish()
//			ctx = nil
//		}
//		scope.Set(jaegerContextKey, ctx)
//		//scope.Set(startTime, time.Now().UnixNano())
//	}
//}
//
//func newAfterCallback() func(*gorm.Scope) {
//	return func(scope *gorm.Scope) {
//		ctx, ok := scopeContext(scope)
//		if !ok {
//			return
//		}
//		span := opentracing.SpanFromContext(ctx)
//		if span == nil {
//			return
//		}
//		defer span.Finish()
//	}
//}
//
//func scopeContext(scope *gorm.Scope) (context.Context, bool) {
//	value, ok := scope.Get(jaegerContextKey)
//	if !ok {
//		return nil, false
//	}
//	ctx, _ := value.(context.Context)
//	return ctx, ctx != nil
//}
