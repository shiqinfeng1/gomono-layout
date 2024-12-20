// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package registrar

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/shiqinfeng1/gomono-layout/pkg/client"
)

// provideFunction1String 为 Function1 提供 string 类型的依赖
func ProvideNacosRegistrarEndpoints(ctx context.Context) []string {
	return g.Cfg().MustGet(ctx, "nacos.endpoints").Strings()
}
func MustNacosRegistrar(endpoints []string) registry.Registrar {
	if len(endpoints) != 1 {
		panic("no such discovery config endpoint")
	}
	c, err := client.NewNacosNamingClient(endpoints[0])
	if err != nil {
		panic(err)
	}
	return nacos.New(c)
}
