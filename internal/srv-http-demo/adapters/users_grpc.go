// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapters

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware/tracing"

	v1 "github.com/shiqinfeng1/gomono-layout/api/gen/user/v1"
	"github.com/shiqinfeng1/gomono-layout/pkg/client"
	"github.com/shiqinfeng1/gomono-layout/pkg/discovery"
	"github.com/shiqinfeng1/gomono-layout/pkg/log"
)

type UserGrpc struct {
	logger    log.Logger
	endpoints []string
	client    v1.UserServiceClient
	close     func() error
}

func NewUserGrpc(endpoints []string) *UserGrpc {
	return &UserGrpc{
		endpoints: endpoints,
		logger: log.WithValues(
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}

func (s UserGrpc) Close() {
	if s.close != nil {
		s.close()
	}
}

func (s *UserGrpc) getClient() v1.UserServiceClient {
	once.Do(func() {
		dis := discovery.MustEtcdDiscovery(s.endpoints)
		conn, err := client.NewGrpcConn(dis, "user")
		if err != nil {
			panic(fmt.Errorf("invalid trainer client from %v: %w", s.endpoints, err))
		}
		s.client = v1.NewUserServiceClient(conn)
		s.close = conn.Close
	})
	return s.client
}

func (s UserGrpc) UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error {
	_, err := s.getClient().UpdateTrainingBalance(ctx, &v1.UpdateTrainingBalanceRequest{
		UserId:       userID,
		AmountChange: int64(amountChange),
	})

	return err
}
