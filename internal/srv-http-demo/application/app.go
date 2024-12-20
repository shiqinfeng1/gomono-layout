// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package application

import (
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/adapters"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/application/command"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/application/query"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/domain/training"
)

// ProviderSet is service providers.

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ApproveTrainingReschedule command.ApproveTrainingRescheduleHandler
	CancelTraining            command.CancelTrainingHandler
	RejectTrainingReschedule  command.RejectTrainingRescheduleHandler
	RescheduleTraining        command.RescheduleTrainingHandler
	RequestTrainingReschedule command.RequestTrainingRescheduleHandler
	ScheduleTraining          command.ScheduleTrainingHandler
}

type Queries struct {
	AllTraining     query.AllTrainingHandler
	TrainingForUser query.TrainingForUserHandler
}

func NewApplication(
	repo training.Repository,
	trainerGrpc *adapters.TrainerGrpc,
	userGrpc *adapters.UserGrpc,
) Application {
	return newApplication(repo, trainerGrpc, userGrpc)
}

// 用于组件测试
func NewComponentTestApplication(repo training.Repository) Application {
	return newApplication(repo, TrainerServiceMock{}, UserServiceMock{})
}

func newApplication(
	repo training.Repository,
	trainerService command.TrainerService,
	userService command.UserService,
) Application {
	logger := log.WithValues(
		"layer", "app",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	return Application{
		Commands: Commands{
			ApproveTrainingReschedule: command.NewApproveTrainingRescheduleHandler(
				repo,
				userService,
				trainerService,
				logger,
			),
			CancelTraining:            command.NewCancelTrainingHandler(repo, userService, trainerService, logger),
			RejectTrainingReschedule:  command.NewRejectTrainingRescheduleHandler(repo, logger),
			RescheduleTraining:        command.NewRescheduleTrainingHandler(repo, userService, trainerService, logger),
			RequestTrainingReschedule: command.NewRequestTrainingRescheduleHandler(repo, logger),
			ScheduleTraining:          command.NewScheduleTrainingHandler(repo, userService, trainerService, logger),
		},
		// todo: 读写分离后，最好不要使用同一个repo
		Queries: Queries{
			AllTraining:     query.NewAllTrainingHandler(repo, logger),
			TrainingForUser: query.NewTrainingForUserHandler(repo, logger),
		},
	}
}
