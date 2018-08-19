package server

import (
	"github.com/kataras/iris"
)

func (s *TypeServer) Routing() {

	s.App.RegisterView(iris.HTML("./templates", ".html"))


	//s.App.Get("/", func(ctx context.Context) {
	//	ctx.Writef("Music Room Project")
	//})


	s.App.Get("/healthcheck", func(ctx iris.Context) {
		ctx.StatusCode(200)
		ctx.JSON(iris.Map{
			"message": "SERVER UP",
		})
	})

	s.App.Get("/auth/{provider}/callback", func(ctx iris.Context) {

		user, err := CompleteUserAuth(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("%v", err)
			return
		}
		ctx.ViewData("", user)
		if err := ctx.View("user.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	s.App.Get("/auth/{provider}", func(ctx iris.Context) {
		// try to get the user without re-authenticating
		if gothUser, err := CompleteUserAuth(ctx); err == nil {
			ctx.ViewData("", gothUser)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			BeginAuthHandler(ctx)
		}
	})

	s.App.Get("/", func(ctx iris.Context) {

		ctx.ViewData("", s.Providers)

		if err := ctx.View("index.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	s.App.Get("/auth/{provider}", func(ctx iris.Context) {
		// try to get the user without re-authenticating
		if gothUser, err := CompleteUserAuth(ctx); err == nil {
			ctx.ViewData("", gothUser)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			BeginAuthHandler(ctx)
		}
	})

}