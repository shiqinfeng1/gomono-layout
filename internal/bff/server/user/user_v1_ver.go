package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/shiqinfeng1/gomono-layout/api/restful/user/v1"
)

func (c *ControllerV1) IsAutoUpdate(ctx context.Context, req *v1.IsAutoUpdateReq) (res *v1.IsAutoUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
