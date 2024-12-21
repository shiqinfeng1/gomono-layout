// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"context"
	"time"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/domain/training"
	"github.com/shiqinfeng1/gomono-layout/pkg/decorator"
	"github.com/shiqinfeng1/gomono-layout/pkg/log"
)

type RescheduleTraining struct {
	TrainingUUID string
	NewTime      time.Time

	User training.User

	NewNotes string
}

type RescheduleTrainingHandler decorator.CommandHandler[RescheduleTraining]

type rescheduleTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewRescheduleTrainingHandler(
	repo training.Repository,
	userService UserService,
	trainerService TrainerService,
	logger log.Logger,
) RescheduleTrainingHandler {
	if repo == nil {
		panic("nil repo")
	}
	if userService == nil {
		panic("nil userService")
	}
	if trainerService == nil {
		panic("nil trainerService")
	}

	return decorator.ApplyCommandDecorators[RescheduleTraining](
		rescheduleTrainingHandler{repo: repo, userService: userService, trainerService: trainerService},
		logger,
	)
}

func (h rescheduleTrainingHandler) Handle(ctx context.Context, cmd RescheduleTraining) (err error) {
	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			originalTrainingTime := tr.Time()

			if err := tr.UpdateNotes(cmd.NewNotes); err != nil {
				return nil, err
			}

			if err := tr.RescheduleTraining(cmd.NewTime); err != nil {
				return nil, err
			}

			err := h.trainerService.MoveTraining(ctx, cmd.NewTime, originalTrainingTime)
			if err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
