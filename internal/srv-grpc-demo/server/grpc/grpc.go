// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/go-kratos/kratos/v2"

	v1 "github.com/shiqinfeng1/gomono-layout/api/gen/training/v1"
	"github.com/shiqinfeng1/gomono-layout/pkg/client"
	"github.com/shiqinfeng1/gomono-layout/pkg/trace"

	kmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "srv-grpc"    // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	SrvInfo = &trace.SrvInfo{
		ID:      ID,
		Name:    Name,
		Version: Version,
	}
)

func NewSrv(ctx context.Context, regstr registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
		),
		kratos.Registrar(regstr),
		kratos.Endpoint(
			&url.URL{Scheme: "http", Host: g.Cfg().MustGet(ctx, "register.endpoints").String()},
		), //  指定服务地址，该地址会提交给注册中心，如果不指定，那么将注册容器内部地址，导致外部无法访问本服务
	)
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(ctx context.Context, si *trace.SrvInfo, training GrpcService) *grpc.Server {

	trace.New(ctx, si, g.Cfg().MustGet(ctx, "trace.endpoint").String())

	opts := []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			kmetrics.Server(
				kmetrics.WithSeconds(client.MetricsSeconds),
				kmetrics.WithRequests(client.MetricsRequests),
			),
		),
	}
	nw := g.Cfg().MustGet(ctx, "grpc.network").String()
	if nw != "" {
		opts = append(opts, grpc.Network(nw))
	}
	ad := g.Cfg().MustGet(ctx, "grpc.addr").String()
	if ad != "" {
		opts = append(opts, grpc.Address(ad))
	}
	to := g.Cfg().MustGet(ctx, "grpc.timeout").Int()
	if to != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(to)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterTrainingServiceServer(srv, training)
	return srv
}
