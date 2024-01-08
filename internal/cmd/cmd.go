package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"mygf/internal/controller/hello"
	"mygf/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//s.Group("/ai", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewV1(),
			//	)
			//})

			//s.BindHandler("POST:/say", func(req *ghttp.Request) {
			//	req.Response.Writeln("Say Hello")
			//})

			//s.BindObject("/user", user.New()) //多个方法

			//s.BindObjectRest("/user", user.New()) //restful

			//s.Group("/user", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		user.New(),
			//	)
			//})

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/hello", func(group1 *ghttp.RouterGroup) {
					group1.Bind(hello.NewV1())
				})
				group.Group("/user", func(group2 *ghttp.RouterGroup) {
					group2.Bind(user.New())
				})

				//group.Bind(
				//	hello.NewV1(),
				//	user.New(),
				//)
			})

			//s.BindHandler("POST:/user/add", func(req *ghttp.Request) {
			//	req.Response.Writeln("Say Hello")
			//})
			s.Run()
			return nil
		},
	}
)
