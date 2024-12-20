// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"context"
	"time"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/domain/training"
	"github.com/shiqinfeng1/gomono-layout/pkg/decorator"
)

type RequestTrainingReschedule struct {
	TrainingUUID string
	NewTime      time.Time

	User training.User

	NewNotes string
}

type RequestTrainingRescheduleHandler decorator.CommandHandler[RequestTrainingReschedule]

type requestTrainingRescheduleHandler struct {
	repo training.Repository
}

func NewRequestTrainingRescheduleHandler(
	repo training.Repository,
	logger log.Logger,
) RequestTrainingRescheduleHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return decorator.ApplyCommandDecorators[RequestTrainingReschedule](
		requestTrainingRescheduleHandler{repo: repo},
		logger,
	)
}

func (h requestTrainingRescheduleHandler) Handle(ctx context.Context, cmd RequestTrainingReschedule) (err error) {
	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.UpdateNotes(cmd.NewNotes); err != nil {
				return nil, err
			}

			tr.ProposeReschedule(cmd.NewTime, cmd.User.Type())

			return tr, nil
		},
	)
}
