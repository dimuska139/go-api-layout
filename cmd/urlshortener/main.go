package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/di"
	"github.com/dimuska139/urlshortener/pkg/logging"

	"github.com/urfave/cli/v2"
	"os"
)

var Version = "master"

func runApplication(c *cli.Context) error {
	ctx := c.Context

	cfg, err := di.InitConfig(c.String("config"), config.VersionParam(Version))
	if err != nil {
		return fmt.Errorf("initialize config: %w", err)
	}

	logging.InitLogger(logging.LogLevel(cfg.Loglevel))

	app, cleanup, err := di.InitApplication(cfg)
	if err != nil {
		return fmt.Errorf("initialize application: %w", err)
	}

	defer cleanup()

	go func() {
		app.Run(ctx)
	}()

	logging.Info(ctx, fmt.Sprintf("SerpParserAPI started at :%d", cfg.HttpServer.Port))

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGTERM)
	signal.Notify(stopSignal, syscall.SIGINT)

	reloadSignal := make(chan os.Signal, 1)
	signal.Notify(reloadSignal, syscall.SIGUSR1)

	for {
		select {
		case <-stopSignal:
			logging.Info(ctx, "Shutdown started")
			app.Stop(ctx)
			logging.Info(ctx, "Shutdown finished")
			os.Exit(0)

		case <-reloadSignal:
			break
		}
	}
}

func migrate(c *cli.Context) error {
	ctx := c.Context

	cfg, err := di.InitConfig(c.String("config"), config.VersionParam(Version))
	if err != nil {
		return fmt.Errorf("initialize config: %w", err)
	}

	logging.InitLogger(logging.LogLevel(cfg.Loglevel))

	migrator, cleanup, err := di.InitMigrator(cfg)
	if err != nil {
		return fmt.Errorf("initialize migrator: %w", err)
	}

	defer cleanup()

	logging.Info(ctx, "Applying migrations...")

	if err := migrator.Up(); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	logging.Info(ctx, "Applying migrations finished")

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Shortener",
		Usage: "Url shortener service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "./config.yml",
				Usage: "path to the config file",
			},
		},
		Action: runApplication,
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "Apply migrations",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config",
						Value: "./config.yml",
						Usage: "path to the config file",
					},
				},
				Action: migrate,
			},
		},
	}

	logging.InitLogger(logging.LogLevelInfo)

	if err := app.Run(os.Args); err != nil {
		logging.Fatal(context.Background(), "Can't run application: ",
			"err", err.Error())
	}
}
