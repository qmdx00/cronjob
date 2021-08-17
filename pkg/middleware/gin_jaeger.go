package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Opentracing middleware for gin request ...
func Opentracing(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := tracer.StartSpan(
			ctx.Request.RequestURI,
			opentracing.Tag{Key: string(ext.Component), Value: "gin request"},
		)
		defer span.Finish()

		// store context.Context to gin.Context
		ctx.Set("context", opentracing.ContextWithSpan(context.Background(), span))

		ctx.Next()
	}
}
