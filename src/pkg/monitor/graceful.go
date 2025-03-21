package monitor

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/gunzgo2mars/skeleton-project/src/pkg/logger"
)

func GracefulShutdownHttpServer(
	ctx context.Context,
	cancel context.CancelFunc,
	httpServer *echo.Echo,
) {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	newCtx, newCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer newCancel()

	select {
	case <-ctx.Done():

		logger.Info(newCtx, "MonitorGraceful",
			zap.String("type", "server"),
			zap.String("msg", "MonitorGraceful - Terminating: context cancelled"),
		)
	case s := <-sigterm:
		logger.Info(newCtx, "MonitorGraceful",
			zap.String("type", "server"),
			zap.String("msg", fmt.Sprintf("MonitorGraceful - Terminating: via signal %v", s)),
		)
	}

	cancel()

	if httpServer != nil {
		if err := httpServer.Shutdown(newCtx); err != nil {
			logger.Error(
				newCtx,
				"MonitorGraceful",
				zap.String("type", "server"),
				zap.String(
					"msg",
					fmt.Sprintf(
						"MonitorGraceful - Terminating: shutdown http server error %v",
						err,
					),
				),
			)
		} else {
			logger.Info(newCtx, "MonitorGraceful",
				zap.String("type", "server"),
				zap.String("msg", "MonitorGraceful - Terminating: shutdown http server success"),
			)
		}
	}
}
