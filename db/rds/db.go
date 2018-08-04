package rds

import (
	"os"
	"strconv"
	"github.com/go-redis/redis"
	"fmt"
)

type TypeRedis struct {
	RDShost string
	RDSport string
	RDSname string
}

func (r *TypeRedis) NewConn() *redis.Client {
	dbname, err := strconv.Atoi(r.RDSname)
	if err != nil {
		panic(err)
	}
	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.RDShost, r.RDSport),
		Password: "", // no password set
		DB:       dbname,
	})

	_, errConn := connection.Ping().Result()
	if errConn != nil {
		panic(errConn)
	}

	return connection
}

func New() *TypeRedis {
	return &TypeRedis{
		os.Getenv("RDSHOST"),
		os.Getenv("RDSPORT"),
		os.Getenv("RDSNAME"),
	}
}
