package globalshared

import (
	"context"

	"go.uber.org/zap"

	"github.com/Bhinneka/candi/logger"
	"github.com/Bhinneka/candi/tracer"
	"github.com/getsentry/sentry-go"
)

// LogParam : logging parameter
type LogParam struct {
	Error                         error
	Message, OperationName, Scope string
	IsSentry                      bool
}

// Log : report log
func Log(ctx context.Context, param LogParam) {
	var level = zap.InfoLevel
	if param.Error != nil {
		level = zap.ErrorLevel
		tracer.SetError(ctx, param.Error)
		if param.IsSentry {
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetTag("traceId", tracer.GetTraceID(ctx))
				sentry.CaptureException(param.Error)
			})
		}
	}
	logger.Log(level, param.Message, param.OperationName, param.Scope)
}
