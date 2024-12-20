// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package query

import (
	"context"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/domain/model"
	"github.com/shiqinfeng1/gomono-layout/pkg/decorator"
)

type TrainingForUser struct {
	uuid string
}

type TrainingForUserHandler decorator.QueryHandler[TrainingForUser, []model.Training]

type trainingForUserHandler struct {
	readModel TrainingForUserReadModel
}

func NewTrainingForUserHandler(
	readModel TrainingForUserReadModel,
	logger log.Logger,
) TrainingForUserHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[TrainingForUser, []model.Training](
		trainingForUserHandler{readModel: readModel},
		logger,
	)
}

type TrainingForUserReadModel interface {
	FindTrainingForUser(ctx context.Context, userUUID string) ([]model.Training, error)
}

func (h trainingForUserHandler) Handle(ctx context.Context, query TrainingForUser) (tr []model.Training, err error) {
	return h.readModel.FindTrainingForUser(ctx, query.uuid)
}
