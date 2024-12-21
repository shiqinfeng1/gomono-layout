// Copyright 2024 Shiqinfeng &lt; 150627601@qq.com >. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	v1 "github.com/shiqinfeng1/gomono-layout/api/restful/user/v1"
)

type IUserV1 interface {
	IsAutoUpdate(ctx context.Context, req *v1.IsAutoUpdateReq) (res *v1.IsAutoUpdateRes, err error)
	GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error)
}
