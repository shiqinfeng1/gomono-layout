// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"context"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/domain/training"
	"github.com/shiqinfeng1/gomono-layout/pkg/decorator"
)

type RejectTrainingReschedule struct {
	TrainingUUID string
	User         training.User
}

type RejectTrainingRescheduleHandler decorator.CommandHandler[RejectTrainingReschedule]

type rejectTrainingRescheduleHandler struct {
	repo training.Repository
}

func NewRejectTrainingRescheduleHandler(
	repo training.Repository,
	logger log.Logger,
) RejectTrainingRescheduleHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return decorator.ApplyCommandDecorators[RejectTrainingReschedule](
		rejectTrainingRescheduleHandler{repo: repo},
		logger,
	)
}

func (h rejectTrainingRescheduleHandler) Handle(ctx context.Context, cmd RejectTrainingReschedule) (err error) {
	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.RejectReschedule(); err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
