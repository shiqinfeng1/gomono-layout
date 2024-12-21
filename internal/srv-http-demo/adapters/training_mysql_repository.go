// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapters

import (
	"context"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/domain/model"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/domain/training"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"
)

type trainingRepo struct {
	log log.Logger
}

// NewTrainingRepo .
func NewTrainingRepo() training.Repository {
	return &trainingRepo{
		log: log.WithValues("repo", "training"),
	}
}

func (r trainingRepo) AddTraining(ctx context.Context, tr *training.Training) error {
	return nil
}

func (r trainingRepo) GetTraining(
	ctx context.Context,
	trainingUUID string,
	user training.User,
) (*training.Training, error) {
	return nil, nil
}

func (r trainingRepo) UpdateTraining(
	ctx context.Context,
	trainingUUID string,
	user training.User,
	updateFn func(ctx context.Context, tr *training.Training) (*training.Training, error),
) error {
	return nil
}

func (r trainingRepo) AllTraining(ctx context.Context) ([]model.Training, error) {
	return nil, nil
}

func (r trainingRepo) FindTrainingForUser(ctx context.Context, userUUID string) ([]model.Training, error) {
	return nil, nil
}

// warning: RemoveAllTraining was designed for tests for doing data cleanups
func (r trainingRepo) RemoveAllTraining(ctx context.Context) error {
	return nil
}
