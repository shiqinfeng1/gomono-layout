// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	trainerv1 "github.com/shiqinfeng1/gomono-layout/api/gen/trainer/v1"
	trainingv1 "github.com/shiqinfeng1/gomono-layout/api/gen/training/v1"
	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/application"
)

type HttpService struct {
	app application.Application
}

func NewService(application application.Application) *HttpService {
	return &HttpService{
		app: application,
	}
}

func (h HttpService) GetTrainerAvailableHours(
	ctx context.Context,
	req *trainerv1.GetTrainerAvailableHoursRequest,
) (*trainerv1.GetTrainerAvailableHoursRespone, error) {
	return &trainerv1.GetTrainerAvailableHoursRespone{}, nil
}

func (h HttpService) MakeHourAvailable(
	ctx context.Context,
	req *trainerv1.MakeHourAvailableRequest,
) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func (h HttpService) MakeHourUnavailable(
	ctx context.Context,
	req *trainerv1.MakeHourUnavailableRequest,
) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func (h HttpService) GetTraining(ctx context.Context, epy *emptypb.Empty) (*trainingv1.GetTrainingResponse, error) {
	return &trainingv1.GetTrainingResponse{}, nil
}

func (h HttpService) CreateTraining(
	ctx context.Context,
	req *trainingv1.CreateTrainingRequest,
) (*trainingv1.CreateTrainingResponse, error) {
	return &trainingv1.CreateTrainingResponse{}, nil
}

func (h HttpService) CancelTraining(
	ctx context.Context,
	req *trainingv1.CancelTrainingRequest,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RescheduleTraining(
	ctx context.Context,
	req *trainingv1.RescheduleTrainingRequest,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) ApproveRescheduleTraining(
	ctx context.Context,
	req *trainingv1.ApproveRescheduleTrainingRequest,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RequestRescheduleTraining(
	ctx context.Context,
	req *trainingv1.RequestRescheduleTrainingRequest,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RejectRescheduleTraining(
	ctx context.Context,
	req *trainingv1.RejectRescheduleTrainingRequest,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
