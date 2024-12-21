// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/fluent/fluent-logger-golang/fluent"
)

type fluentWriteSyncer struct {
	output *fluent.Fluent
}

func NewFluentWriteSyncer(addr string) (*fluentWriteSyncer, error) {
	c := fluent.Config{
		Async:              true,
		ForceStopAsyncSend: true,
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "tcp":
		host, port, err2 := net.SplitHostPort(u.Host)
		if err2 != nil {
			return nil, err2
		}
		if c.FluentPort, err = strconv.Atoi(port); err != nil {
			return nil, err
		}
		c.FluentNetwork = u.Scheme
		c.FluentHost = host
	case "unix":
		c.FluentNetwork = u.Scheme
		c.FluentSocketPath = u.Path
	default:
		return nil, fmt.Errorf("unknown network: %s", u.Scheme)
	}
	// 检查地址是否存在服务
	_, err = net.Dial(u.Scheme, u.Host)
	if err != nil {
		return nil, err
	}
	fl, err := fluent.New(c)
	if err != nil {
		return nil, err
	}
	return &fluentWriteSyncer{
		output: fl,
	}, nil
}

func (x *fluentWriteSyncer) Write(data []byte) (n int, err error) {
	if err := x.output.Post("zerolog", data); err != nil {
		return 0, err
	}
	return len(data), nil
}
