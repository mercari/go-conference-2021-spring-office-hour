package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mercari/mercari-microservices-example/pkg/logger"
	"github.com/mercari/mercari-microservices-example/services/authority/grpc"
	"golang.org/x/sys/unix"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			// Unhandleable, something went wrong...
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		return 1
	}
	alogger := l.WithName("authority")

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, 5000, alogger.WithName("grpc"))
	}()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
		return 1
	case <-ctx.Done():
		fmt.Println("shutting down...")
		return 0
	}
}
