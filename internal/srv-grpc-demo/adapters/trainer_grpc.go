// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	v1 "github.com/shiqinfeng1/gomono-layout/api/gen/trainer/v1"
	"github.com/shiqinfeng1/gomono-layout/pkg/client"
	"github.com/shiqinfeng1/gomono-layout/pkg/discovery"
)

var once sync.Once

type TrainerGrpc struct {
	logger    log.Logger
	endpoints []string
	client    v1.TrainerServiceClient
	close     func() error
}

func NewTrainerGrpc(endpoints []string) *TrainerGrpc {
	return &TrainerGrpc{
		endpoints: endpoints,
		logger: log.WithValues(
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}

func (s TrainerGrpc) Close() {
	if s.close != nil {
		s.close()
	}
}

func (s *TrainerGrpc) trainerServiceClient() v1.TrainerServiceClient {
	once.Do(func() {
		dis := discovery.MustEtcdDiscovery(s.endpoints)
		conn, err := client.NewGrpcConn(dis, "trainer")
		if err != nil {
			panic(fmt.Errorf("invalid trainer client from %v: %w", s.endpoints, err))
		}
		s.client = v1.NewTrainerServiceClient(conn)
		s.close = conn.Close
	})
	return s.client
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.trainerServiceClient().ScheduleTraining(ctx, &v1.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.trainerServiceClient().CancelTraining(ctx, &v1.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) MoveTraining(
	ctx context.Context,
	newTime time.Time,
	originalTrainingTime time.Time,
) error {
	err := s.ScheduleTraining(ctx, newTime)
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}

	err = s.CancelTraining(ctx, originalTrainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to cancel training")
	}

	return nil
}
