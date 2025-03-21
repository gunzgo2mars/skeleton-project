package main

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/gunzgo2mars/skeleton-project/src/internal/router"
	"github.com/gunzgo2mars/skeleton-project/src/pkg/configurer"
	"github.com/gunzgo2mars/skeleton-project/src/pkg/logger"
	"github.com/gunzgo2mars/skeleton-project/src/pkg/monitor"
)

func main() {
	// initialzing context with cancel.
	ctx, cancel := context.WithCancel(context.Background())

	// initialzing load .yaml and dotenv configuration.
	var conf *configurer.AppConfig
	if err := configurer.LoadConfig(&conf, "config", "config", "yaml", "APPENV"); err != nil {
		panic(err.Error())
	}

	if err := configurer.LoadDotEnv(&conf.Secrets, "./config/", "SECRET", "APPENV"); err != nil {
		panic(err.Error())
	}

	// initialzing logger.
	logger.InitLogger(conf.Log.Env)
	defer logger.Sync()

	// initialzing dependencies.

	// initialzing repositories.

	// initialzing services.

	// initialzing handlers.

	// initialzing httpserver.
	httpServer := router.New(conf.App.Port)
	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Fatal(
				ctx,
				fmt.Sprintf("Error: %s", err.Error()),
				zap.String("type", "exec"),
			)
		}
	}()

	monitor.GracefulShutdownHttpServer(ctx, cancel, httpServer.Server())
}
