// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/domain/training"
	"github.com/shiqinfeng1/gomono-layout/pkg/decorator"
)

type ScheduleTraining struct {
	TrainingUUID string

	UserUUID string
	UserName string

	TrainingTime time.Time
	Notes        string
}

type ScheduleTrainingHandler decorator.CommandHandler[ScheduleTraining]

type scheduleTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewScheduleTrainingHandler(
	repo training.Repository,
	userService UserService,
	trainerService TrainerService,
	logger log.Logger,
) ScheduleTrainingHandler {
	if repo == nil {
		panic("nil repo")
	}
	if userService == nil {
		panic("nil repo")
	}
	if trainerService == nil {
		panic("nil trainerService")
	}

	return decorator.ApplyCommandDecorators[ScheduleTraining](
		scheduleTrainingHandler{repo: repo, userService: userService, trainerService: trainerService},
		logger,
	)
}

func (h scheduleTrainingHandler) Handle(ctx context.Context, cmd ScheduleTraining) (err error) {
	tr, err := training.NewTraining(cmd.TrainingUUID, cmd.UserUUID, cmd.UserName, cmd.TrainingTime)
	if err != nil {
		return err
	}

	if err := h.repo.AddTraining(ctx, tr); err != nil {
		return err
	}

	err = h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), -1)
	if err != nil {
		return errors.Wrap(err, "unable to change training balance")
	}

	err = h.trainerService.ScheduleTraining(ctx, tr.Time())
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}

	return nil
}
