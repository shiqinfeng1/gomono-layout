// Copyright 2024 Shiqinfeng &lt; 150627601@qq.com >. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/shiqinfeng1/gomono-layout/internal/srv-http-demo/server/http"
)

var srv = &gcmd.Command{
	Name:        "svc",
	Brief:       "start training grpc service",
	Description: "entry to start a training grpc service",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		srv, cleanup, err := wireSrv(
			ctx,
			http.SrvInfo,
		)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := srv.Run(); err != nil {
			panic(err)
		}
		return nil
	},
}

var root = gcmd.Command{
	Name:  "main",
	Usage: "main <sub-command>",
	Brief: "this is main command, please specify a sub command",
}

func New() gcmd.Command {
	err := root.AddCommand(srv)
	if err != nil {
		panic(err)
	}
	return root
}
