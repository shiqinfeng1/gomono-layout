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

type AllTraining struct{}

type AllTrainingHandler decorator.QueryHandler[AllTraining, []model.Training]

type allTrainingHandler struct {
	readModel AllTrainingReadModel
}

func NewAllTrainingHandler(
	readModel AllTrainingReadModel,
	logger log.Logger,
) AllTrainingHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[AllTraining, []model.Training](
		allTrainingHandler{readModel: readModel},
		logger,
	)
}

type AllTrainingReadModel interface {
	AllTraining(ctx context.Context) ([]model.Training, error)
}

func (h allTrainingHandler) Handle(ctx context.Context, _ AllTraining) (tr []model.Training, err error) {
	return h.readModel.AllTraining(ctx)
}
