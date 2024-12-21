// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package training

import (
	"context"
	"fmt"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-grpc-demo/domain/model"
)

// 公共错误，所有实现Repository接口的实例，都会用到该错误
type NotFoundError struct {
	TrainingUUID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("training '%s' not found", e.TrainingUUID)
}

// 定义repo接口，在领域层中定义实现依赖倒置
type Repository interface {
	AddTraining(ctx context.Context, tr *Training) error

	GetTraining(ctx context.Context, trainingUUID string, user User) (*Training, error)

	UpdateTraining(
		ctx context.Context,
		trainingUUID string,
		user User,
		updateFn func(ctx context.Context, tr *Training) (*Training, error),
	) error

	AllTraining(ctx context.Context) ([]model.Training, error)
	FindTrainingForUser(ctx context.Context, userUUID string) ([]model.Training, error)
}
