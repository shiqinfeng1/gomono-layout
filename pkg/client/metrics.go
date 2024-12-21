// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	// "github.com/prometheus/client_golang/prometheus"
	kmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var (
	meter           = sdkmetric.NewMeterProvider().Meter("name")
	MetricsSeconds  metric.Float64Histogram
	MetricsRequests metric.Int64Counter
)

func init() {
	MetricsSeconds, _ = kmetrics.DefaultSecondsHistogram(meter, "histogramName")
	MetricsRequests, _ = kmetrics.DefaultRequestsCounter(meter, "counterName")

}

// var (
// 	// 设置 metrics 中间件统计请求耗时的 Observer 直方图
// 	MetricsSeconds = prometheus.NewHistogram(prometheus.HistogramOpts{
// 		Namespace: "server",
// 		Subsystem: "requests",
// 		Name:      "duration_sec",
// 		Help:      "server requests duration(sec).",
// 		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
// 	})

// 	// 设置 metrics 中间件统计请求计数的 Counter 计数器
// 	MetricsRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Namespace: "client",
// 		Subsystem: "requests",
// 		Name:      "code_total",
// 		Help:      "The total number of processed requests",
// 	}, []string{"kind", "operation", "code", "reason"})
// 	//  统计系统资源
// 	MetricsLoads = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Namespace: "server",
// 		Subsystem: "system",
// 		Name:      "load_total",
// 		Help:      "The load of cpu & memory",
// 	}, []string{"kind", "operation", "code", "reason"})
// )

// func init() {
// 	prometheus.MustRegister(MetricsSeconds, MetricsRequests, MetricsLoads)
// }
