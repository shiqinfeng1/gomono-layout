// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package discovery

import (
	"errors"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/shiqinfeng1/gomono-layout/pkg/client"
)

var ErrServiceKindInvaild = errors.New("service discovery kind is invalid")

func MustNacosDiscovery(endpoint string, kind string) registry.Discovery {
	if kind != "grpc" && kind != "http" {
		panic(ErrServiceKindInvaild)
	}
	c, err := client.NewNacosNamingClient(endpoint)
	if err != nil {
		panic(err)
	}

	return nacos.New(c, nacos.WithDefaultKind(kind))
}
