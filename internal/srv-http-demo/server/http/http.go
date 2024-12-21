// Copyright 2024 Shiqinfeng &lt; 150627601@qq.com >. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/handlers"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/shiqinfeng1/gomono-layout/pkg/client"
	"github.com/shiqinfeng1/gomono-layout/pkg/trace"

	"net/url"
	"os"

	"github.com/go-kratos/kratos/v2"
	kjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	kmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/golang-jwt/jwt/v5"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "srv-http"    // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	SrvInfo = &trace.SrvInfo{
		ID:      ID,
		Name:    Name,
		Version: Version,
	}
)

func NewSrv(ctx context.Context, regstr registry.Registrar, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
		),
		kratos.Registrar(regstr),
		kratos.Endpoint(
			&url.URL{Scheme: "http", Host: g.Cfg().MustGet(ctx, "register.endpoints").String()},
		), //  指定服务地址，该地址会提交给注册中心，如果不指定，那么将注册容器内部地址，导致外部无法访问本服务
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(ctx context.Context, si *trace.SrvInfo, s *HttpService) *http.Server {
	trace.New(ctx, si, g.Cfg().MustGet(ctx, "trace.endpoint").String())
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			selector.Server(
				kjwt.Server(
					func(token *jwt.Token) (interface{}, error) {
						return []byte(g.Cfg().MustGet(ctx, "jwt.apikey").String()), nil
					},
					kjwt.WithSigningMethod(jwt.SigningMethodHS256),
					kjwt.WithClaims(func() jwt.Claims {
						return &jwt.MapClaims{}
					})),
			).Build(),
			kmetrics.Server(
				kmetrics.WithSeconds(client.MetricsSeconds),
				kmetrics.WithRequests(client.MetricsRequests),
			),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	nw := g.Cfg().MustGet(ctx, "http.network").String()
	if nw != "" {
		opts = append(opts, http.Network(nw))
	}
	ad := g.Cfg().MustGet(ctx, "grpc.addr").String()
	if ad != "" {
		opts = append(opts, http.Address(ad))
	}
	to := g.Cfg().MustGet(ctx, "grpc.timeout").Int()
	if to != 0 {
		opts = append(opts, http.Timeout(time.Duration(to)*time.Second))
	}
	srv := http.NewServer(opts...)
	return srv
}
