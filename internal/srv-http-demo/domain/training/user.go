// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package training

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrIncorrectInput = func(userType string) error { return errors.New(fmt.Sprintf("invalid '%s' role", userType)) }

// UserType is enum-like type.
// We are using struct instead of string, to ensure about immutability.
// 注释：使用结构体来表示枚举的类型信息，是为了防止字符串类型强转导致类型改变
// 例如：type UserType string ； aa := UserType("unknowType"); aa并非我们定义的类型， 但能够被当做UserType来使用
type UserType struct {
	s string
}

func (u UserType) IsZero() bool {
	return u == UserType{}
}

func (u UserType) String() string {
	return u.s
}

var (
	Trainer  = UserType{"trainer"}
	Attendee = UserType{"attendee"}
)

// 类型转换
func NewUserTypeFromString(userType string) (UserType, error) {
	switch userType {
	case "trainer":
		return Trainer, nil
	case "attendee":
		return Attendee, nil
	}

	return UserType{}, ErrIncorrectInput(userType)
}

// 该User是training领域内的用户抽象，并且只用于自己领域。
// 所有外部领域的用户，比如其他领域的trainer，attendee，都需要抽象为User结构
type User struct {
	userUUID string
	userType UserType
}

func (u User) UUID() string {
	return u.userUUID
}

func (u User) Type() UserType {
	return u.userType
}

func (u User) IsEmpty() bool {
	return u == User{}
}

func NewUser(userUUID string, userType UserType) (User, error) {
	if userUUID == "" {
		return User{}, errors.New("missing user UUID")
	}
	if userType.IsZero() {
		return User{}, errors.New("missing user type")
	}

	return User{userUUID: userUUID, userType: userType}, nil
}

// 用于测试case
func MustNewUser(userUUID string, userType UserType) User {
	u, err := NewUser(userUUID, userType)
	if err != nil {
		panic(err)
	}

	return u
}

// 就近定义业务错误，该错误需要记录信息，因此使用结构体
type ForbiddenToSeeTrainingError struct {
	RequestingUserUUID string
	TrainingOwnerUUID  string
}

// 实现error的Error方法
func (f ForbiddenToSeeTrainingError) Error() string {
	return fmt.Sprintf(
		"user '%s' can't see user '%s' training",
		f.RequestingUserUUID, f.TrainingOwnerUUID,
	)
}

// 暂时不太清楚为啥单独封装一个函数，而不是作为User的一个接口函数
// 检查用户是否能看到训练信息。如果是教练：能看到全部；如果是学员：能看到自己相关的；
func CanUserSeeTraining(user User, training Training) error {
	if user.Type() == Trainer {
		return nil
	}
	if user.UUID() == training.UserUUID() {
		return nil
	}

	return ForbiddenToSeeTrainingError{user.UUID(), training.UserUUID()}
}
