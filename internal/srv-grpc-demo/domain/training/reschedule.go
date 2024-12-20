// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package training

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func (t Training) MovedProposedBy() UserType {
	return t.moveProposedBy
}

func (t Training) ProposedNewTime() time.Time {
	return t.proposedNewTime
}

type CantRescheduleBeforeTimeError struct {
	TrainingTime time.Time
}

func (c CantRescheduleBeforeTimeError) Error() string {
	return fmt.Sprintf(
		"can't reschedule training, not enough time before, training time: %s",
		c.TrainingTime,
	)
}

// 直接重新安排
func (t *Training) RescheduleTraining(newTime time.Time) error {
	if !t.CanBeCanceledForFree() {
		err := CantRescheduleBeforeTimeError{
			TrainingTime: t.Time(),
		}
		return errors.WithStack(err)
	}

	t.time = newTime

	return nil
}

// 申请安排重新
func (t *Training) ProposeReschedule(newTime time.Time, proposerType UserType) {
	t.moveProposedBy = proposerType
	t.proposedNewTime = newTime
}

func (t *Training) IsRescheduleProposed() bool {
	return !t.moveProposedBy.IsZero() && !t.proposedNewTime.IsZero()
}

var ErrNoRescheduleRequested = errors.New("no training reschedule was requested yet")

// 批准重新安排
func (t *Training) ApproveReschedule(userType UserType) error {
	if !t.IsRescheduleProposed() {
		return errors.WithStack(ErrNoRescheduleRequested)
	}

	if t.moveProposedBy == userType {
		return errors.Errorf(
			"trying to approve reschedule by the same user type which proposed reschedule (%s)",
			userType.String(),
		)
	}

	t.time = t.proposedNewTime

	t.proposedNewTime = time.Time{}
	t.moveProposedBy = UserType{}

	return nil
}

// 拒绝重新安排
func (t *Training) RejectReschedule() error {
	if !t.IsRescheduleProposed() {
		return errors.WithStack(ErrNoRescheduleRequested)
	}

	t.proposedNewTime = time.Time{}
	t.moveProposedBy = UserType{}

	return nil
}
