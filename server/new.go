package server

import (
	"github.com/music-room/api-server/db/psql"
	"github.com/music-room/api-server/db/rds"
	"github.com/kataras/iris"
)

type TypeServer struct {
	PSQL  *psql.TypePSQL
	Redis *rds.TypeRedis
	App   *iris.Application
}

func New() *TypeServer {
	return &TypeServer{
		PSQL:  psql.New(),
		Redis: rds.New(),
		App:   iris.New(),
	}
}
