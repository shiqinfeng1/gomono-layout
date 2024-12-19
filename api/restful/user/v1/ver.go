// Copyright 2024 Shiqinfeng &lt; 150627601@qq.com >. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import "github.com/gogf/gf/v2/frame/g"

type IsAutoUpdateReq struct {
	g.Meta `     path:"/version/isAutoupdate" tags:"版本管理" method:"post" summary:"配置是否自动更新"`
	IsAuto bool `                                                                          p:"is_auto" dc:"ture:自动更新"`
}
type IsAutoUpdateRes struct {
	g.Meta `mime:"application/json"`
}
type GetConfigReq struct {
	g.Meta `path:"/version/configures" tags:"版本管理" method:"post" summary:"查询是否自动更新"`
}
type GetConfigRes struct {
	g.Meta     `       mime:"application/json"`
	IsAuto     bool   `                        json:"is_auto"     dc:"ture:自动更新"`
	ExeVersion string `                        json:"exe_version" dc:"本exe的版本号"`
	VerCommit  string `                        json:"ver_commit"  dc:"版本描述"`
}
