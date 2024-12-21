// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package decorator

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"
)

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger log.Logger) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: handler,
		logger: logger.WithValues(
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
