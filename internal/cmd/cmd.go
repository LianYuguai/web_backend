package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"web_backend/internal/controller"
	"web_backend/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server("0.0.0.0")
			s.SetPort(8888)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				group.Bind(
					controller.Hello,
				)
				group.Bind(
					controller.User,
				)
				group.Bind(
					controller.Dongdian,
				)
			})
			s.Run()
			return nil
		},
	}
)
