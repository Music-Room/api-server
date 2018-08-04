package main

import (
	"context"
	"github.com/music-room/api-server/server"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/kataras/iris"
)

func main() {

	apiServer := server.New()

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("Shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			apiServer.App.Shutdown(ctx)
		}
	}()

	apiServer.App.Logger().SetLevel("debug")

	apiServer.Routing()

	// Start the server and disable the default interrupt handler in order to
	errServer := apiServer.App.Run(
		// Start the web server at localhost:80
		iris.Addr(":9000"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
		// handle it clear and simple by our own, without any issues.
		iris.WithoutInterruptHandler)

	if errServer != nil {
		panic(errServer)
	}
}
