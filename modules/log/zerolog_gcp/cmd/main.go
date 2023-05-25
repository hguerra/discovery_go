package main

import (
	"context"

	"github.com/hguerra/discovery_go/modules/log/zerolog_gcp/pkg/logging"
	"github.com/hirosassa/zerodriver"
)

func main() {
	l1 := zerodriver.NewDevelopmentLogger()
	l1.Info().Str("key", "value").Msg("hello world")
	l1.Info().Labels(zerodriver.Label("foo", "var")).Msg("labeled log")

	traceId := "123"
	spanId := "456"
	l1.Info().TraceContext(traceId, spanId, true, "my-project").Msg("trace contexts")

	l2 := logging.NewLogger("MyService")
	l2.Info().Msg("test")

	ctx := context.WithValue(context.Background(), "userID", "1234")
	l3 := logging.New("gcp-project", "MyService")
	l3.Infof(ctx, "Oi %s", "heitor")
	l3.Debugf(ctx, "debug")
	l3.Infof(ctx, "info")
	l3.Warnf(ctx, "warn")
	l3.Errorf(ctx, "error")
	// l3.Panicf(ctx, "panic")
}
