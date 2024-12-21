// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package training

import (
	"time"

	"github.com/pkg/errors"
)

// 领域层的核心数据结构
type Training struct {
	uuid string

	userUUID string
	userName string

	time  time.Time
	notes string

	proposedNewTime time.Time
	moveProposedBy  UserType

	canceled bool
}

// NewTraining 实例化一个领域数据，每个参数都是必填
func NewTraining(uuid string, userUUID string, userName string, trainingTime time.Time) (*Training, error) {
	if uuid == "" {
		return nil, errors.New("empty training uuid")
	}
	if userUUID == "" {
		return nil, errors.New("empty userUUID")
	}
	if userName == "" {
		return nil, errors.New("empty userName")
	}
	if trainingTime.IsZero() {
		return nil, errors.New("zero training time")
	}

	return &Training{
		uuid:     uuid,
		userUUID: userUUID,
		userName: userName,
		time:     trainingTime,
	}, nil
}

// UnmarshalTrainingFromDatabase 转换数据库实体为领域实体
// 注意：该函数只能用于从数据库转换数据为领域实体，而不能用于构造领域实体
func UnmarshalTrainingFromDatabase(
	uuid string,
	userUUID string,
	userName string,
	trainingTime time.Time,
	notes string,
	canceled bool,
	proposedNewTime time.Time,
	moveProposedBy UserType,
) (*Training, error) {
	tr, err := NewTraining(uuid, userUUID, userName, trainingTime)
	if err != nil {
		return nil, err
	}

	tr.notes = notes
	tr.proposedNewTime = proposedNewTime
	tr.moveProposedBy = moveProposedBy
	tr.canceled = canceled

	return tr, nil
}

// 领域实体字段访问的接口封装
func (t Training) UUID() string {
	return t.uuid
}

// 领域实体字段访问的接口封装
func (t Training) UserUUID() string {
	return t.userUUID
}

// 领域实体字段访问的接口封装
func (t Training) UserName() string {
	return t.userName
}

// 领域实体字段访问的接口封装
func (t Training) Time() time.Time {
	return t.time
}

var ErrNoteTooLong = errors.New("Note too long")

func (t *Training) UpdateNotes(notes string) error {
	if len(notes) > 1000 {
		return errors.WithStack(ErrNoteTooLong)
	}

	t.notes = notes
	return nil
}

// 领域实体字段访问的接口封装
func (t Training) Notes() string {
	return t.notes
}
