// Copyright 2024 Shiqinfeng &lt;150627601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "go.uber.org/automaxprocs"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/server"
)

func main() {
	s := server.New()
	s.Run(gctx.GetInitCtx())
}
