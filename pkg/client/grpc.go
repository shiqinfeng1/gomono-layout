// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	ggrpc "google.golang.org/grpc"
)

// 微服务内部之间的通信，无需安全选项
func NewGrpcConn(discovery registry.Discovery, srvName string) (*ggrpc.ClientConn, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+srvName),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
