// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package decorator

import (
	"context"
	"fmt"

	"github.com/shiqinfeng1/gomono-layout/pkg/log"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C] // 下一层命令
	logger log.Logger        // 当前包装的参数
}

// 每层装饰都要实现Handle接口，除了做该层的事之外，还负责调用下一层
func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	// 执行前log
	handlerType := generateActionName(cmd)
	logger := d.logger.WithValues("command", handlerType, "command_body", fmt.Sprintf("%#v", cmd))
	logger.Debug("Executing command")
	// 执行后log
	defer func() {
		if err == nil {
			logger.Info("command executed successfully")
		} else {
			logger.Errorf("failed to execute command: %v", err)
		}
	}()
	// 调用下一层
	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R] // 下一层命令
	logger log.Logger         // 当前包装的参数
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	// 执行前log
	handlerType := generateActionName(cmd)
	logger := d.logger.WithValues("query", handlerType, "query_body", fmt.Sprintf("%#v", cmd))
	logger.Debug("Executing query")
	// 执行后log
	defer func() {
		if err == nil {
			logger.Info("query executed successfully")
		} else {
			logger.Errorf("failed to execute query:%v", err)
		}
	}()
	// 调用下一层
	return d.base.Handle(ctx, cmd)
}
