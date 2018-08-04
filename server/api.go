package server

import (
	"github.com/kataras/iris"
)


func (s *TypeServer) Routing() {

	s.App.Get("/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from Music-Room SERVER")
	})

	s.App.PartyFunc("/room", func(r iris.Party) {
		r.Get("/welcome", func(ctx iris.Context) {
			ctx.Writef("Hello from ROOM")
		})
	})


}