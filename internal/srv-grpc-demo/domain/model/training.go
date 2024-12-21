// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import "time"

type Training struct {
	UUID     string
	UserUUID string
	User     string

	Time  time.Time
	Notes string

	ProposedTime   *time.Time
	MoveProposedBy *string

	CanBeCancelled bool
}
