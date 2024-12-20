// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"context"
	"time"
)

type UserService interface {
	UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error
}

type TrainerService interface {
	ScheduleTraining(ctx context.Context, trainingTime time.Time) error
	CancelTraining(ctx context.Context, trainingTime time.Time) error

	MoveTraining(
		ctx context.Context,
		newTime time.Time,
		originalTrainingTime time.Time,
	) error
}
