// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package training

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func (t Training) CanBeCanceledForFree() bool {
	return time.Until(t.time) >= time.Hour*24
}

// 就近定义错误
var ErrTrainingAlreadyCanceled = errors.New("training is already canceled")

// 取消训练
func (t *Training) Cancel() error {
	if t.IsCanceled() {
		return ErrTrainingAlreadyCanceled // 任何错误不要直接返回 errors.New(...)
	}

	t.canceled = true
	return nil
}

func (t Training) IsCanceled() bool {
	return t.canceled
}

// 暂时不太清楚为啥单独封装一个函数，而不是作为Training的一个接口函数
// CancelBalanceDelta return training balance delta that should be adjusted after training cancellation.
func CancelBalanceDelta(tr Training, cancelingUserType UserType) int {
	if tr.CanBeCanceledForFree() {
		// just give training back
		return 1
	}

	switch cancelingUserType {
	case Trainer:
		// 1 for cancelled training +1 "fine" for cancelling by trainer less than 24h before training
		return 2
	case Attendee:
		// "fine" for cancelling less than 24h before training
		return 0
	default:
		panic(fmt.Sprintf("not supported user type %s", cancelingUserType))
	}
}
