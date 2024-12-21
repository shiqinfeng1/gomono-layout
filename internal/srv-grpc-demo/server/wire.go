// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/adapters"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/application"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/server/grpc"
	"github.com/shiqinfeng1/gomono-layout/pkg/registrar"
	"github.com/shiqinfeng1/gomono-layout/pkg/trace"
)

// wireApp init kratos application.
func wireSrv(
	ctx context.Context,
	srvInfo *trace.SrvInfo,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		adapters.ProviderSet,
		application.ProviderSet,
		registrar.ProviderSet,
		grpc.ProviderSet,
		grpc.NewSrv))
}
