// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package application

import (
	"context"
	"time"
)

type TrainerServiceMock struct {
}

func (t TrainerServiceMock) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	return nil
}

func (t TrainerServiceMock) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	return nil
}

func (t TrainerServiceMock) MoveTraining(ctx context.Context, newTime time.Time, originalTrainingTime time.Time) error {
	return nil
}

type UserServiceMock struct {
}

func (u UserServiceMock) UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error {
	return nil
}
