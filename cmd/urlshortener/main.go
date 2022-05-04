package main

import (
	"fmt"
	"github.com/dimuska139/urlshortener/internal/wire"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

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
		Action: func(c *cli.Context) error {
			cfg, err := wire.InitConfig(c.String("config"))
			logger := wire.InitLogger(cfg)

			if err != nil {
				logger.Error("initialize config service failed", err, map[string]interface{}{})
				return err
			}

			restAPI, err := wire.InitRestAPI(cfg, logger)
			if err != nil {
				logger.Error("initialize REST API failed", err, map[string]interface{}{})
				return err
			}

			go func() {
				if err := restAPI.Start(); err != nil {
					logger.Fatal(fmt.Errorf("serve API service failed: %w", err), map[string]interface{}{})
				}
				logger.Info(fmt.Sprintf("HTTP server started (port: %d)", cfg.Port), nil, map[string]interface{}{})
			}()

			stopSignal := make(chan os.Signal)
			signal.Notify(stopSignal, syscall.SIGTERM)
			signal.Notify(stopSignal, syscall.SIGINT)
			signal.Notify(stopSignal, syscall.SIGKILL)

			reloadSignal := make(chan os.Signal)
			signal.Notify(reloadSignal, syscall.SIGUSR1)
			logger.Info("markup service started", nil, map[string]interface{}{})
			for {
				select {
				case <-stopSignal:
					logger.Info("shutdown started", nil, map[string]interface{}{})
					logger.Info("shutdown finished", nil, map[string]interface{}{})
					os.Exit(0)

				case <-reloadSignal:
					break
				}
			}

			return nil
		},
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
				Action: func(c *cli.Context) error {
					cfg, err := wire.InitConfig(c.String("config"))
					logger := wire.InitLogger(cfg)

					if err != nil {
						logger.Error("initialize config service failed", err, map[string]interface{}{})
						return err
					}

					migrator, err := wire.InitMigrator(cfg, logger)
					if err != nil {
						logger.Error("initialize REST API failed", err, map[string]interface{}{})
						return err
					}

					logger.Info("applying migrations", nil, nil)
					if err := migrator.Up(); err != nil {
						logger.Error("unable apply migrations", err, nil)
						return err
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger := wire.InitLogger(nil)
		logger.Fatal(fmt.Errorf("unable to run application: %w", err), nil)
	}
}
