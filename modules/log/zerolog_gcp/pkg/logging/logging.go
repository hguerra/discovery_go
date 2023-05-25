package logging

import (
	"context"

	"github.com/hirosassa/zerodriver"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

const userIDKey = "userID"

type Logger struct {
	projectID string
	name      string
	sampled   bool
	logger    *zerodriver.Logger
}

func (l *Logger) withTrace(ctx context.Context, event *zerodriver.Event) *zerolog.Event {
	s := trace.SpanFromContext(ctx)
	return event.TraceContext(
		s.SpanContext().TraceID().String(),
		s.SpanContext().SpanID().String(),
		l.sampled,
		l.projectID,
	)
}

func (l *Logger) msgf(ctx context.Context, event *zerodriver.Event, msg string, v ...any) {
	d := zerolog.Dict().
		Str("name", l.name)

	userID, ok := ctx.Value(userIDKey).(string)
	if ok {
		d.Str(userIDKey, userID)
	}

	l.
		withTrace(ctx, event).
		Dict("logging.googleapis.com/labels", d).
		Msgf(msg, v...)
}

func (l *Logger) Debugf(ctx context.Context, msg string, v ...any) {
	l.msgf(ctx, l.logger.Debug(), msg, v...)
}

func (l *Logger) Infof(ctx context.Context, msg string, v ...any) {
	l.msgf(ctx, l.logger.Info(), msg, v...)
}

func (l *Logger) Warnf(ctx context.Context, msg string, v ...any) {
	l.msgf(ctx, l.logger.Warn(), msg, v...)
}

func (l *Logger) Errorf(ctx context.Context, msg string, v ...any) {
	l.msgf(ctx, l.logger.Error(), msg, v...)
}

func (l *Logger) Panicf(ctx context.Context, msg string, v ...any) {
	l.msgf(ctx, l.logger.Panic(), msg, v...)
}

func logger() *zerodriver.Logger {
	return zerodriver.NewProductionLogger()
}

func New(projectID string, name string) *Logger {
	return &Logger{
		projectID: projectID,
		name:      name,
		sampled:   true,
		logger:    logger(),
	}
}

func NewLogger(name string) zerolog.Logger {
	return logger().With().Str("name", name).Logger()
}

func WithContext(ctx context.Context) context.Context {
	return logger().WithContext(ctx)
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
