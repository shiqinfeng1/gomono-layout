// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package discovery

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/shiqinfeng1/gomono-layout/pkg/client"
)

func MustEtcdDiscovery(endpoints []string) registry.Discovery {
	eclient, err := client.NewEtcd(endpoints)
	if err != nil {
		panic(err)
	}
	return etcd.New(eclient)
}
