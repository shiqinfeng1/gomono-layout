// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package application

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewApplication)
