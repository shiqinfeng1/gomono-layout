// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package registrar

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.  使用哪个就注册哪个
var ProviderSet = wire.NewSet(ProvideNacosRegistrarEndpoints, MustNacosRegistrar)
