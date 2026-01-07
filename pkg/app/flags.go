package app

import (
	"context"
	"io"

	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_error/exporter"
	libraryLogger "github.com/ssst0n3/awesome_libs/log"
	"github.com/urfave/cli/v3"
)

var (
	debugFlag = &cli.BoolFlag{
		Name:  "debug",
		Value: false,
		Usage: "Output information for helping debugging sploit",
	}
	experimentalFlag = &cli.BoolFlag{
		Name:  "experimental",
		Value: false,
		Usage: "enable experimental feature",
	}
	colorfulFlag = &cli.BoolFlag{
		Name:  "colorful",
		Value: false,
		Usage: "output colorfully",
	}
	jsonFlag = &cli.BoolFlag{
		Name:  "json",
		Value: false,
		Usage: "output in json format",
	}
)

func InstallGlobalFlagDebug(app *cli.Command, appLogger *logrus.Logger) {
	app.Flags = append(app.Flags, debugFlag)
	before := app.Before
	app.Before = func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		if before != nil {
			_, err := before(ctx, cmd)
			if err != nil {
				return nil, err
			}
		}
		debug := cmd.Bool("debug")
		awesome_error.Default = exporter.GetAwesomeError(appLogger, debug)
		if !debug {
			// print library's logger only if debug mode
			libraryLogger.Logger.SetOutput(io.Discard)
		} else {
			appLogger.Level = logrus.DebugLevel
			appLogger.SetReportCaller(true)
			appLogger.SetFormatter(&logrus.TextFormatter{
				ForceColors: true,
			})
			appLogger.Debug("debug mode on")

			libraryLogger.Logger.Level = logrus.DebugLevel
			libraryLogger.Logger.SetReportCaller(true)
			libraryLogger.Logger.SetFormatter(&logrus.TextFormatter{
				ForceColors: true,
			})
			libraryLogger.Logger.Debug("debug mode on")
		}
		return nil, nil
	}
}

func InstallGlobalFlagExperimentalFlag(app *cli.Command) {
	app.Flags = append(app.Flags, experimentalFlag)
}

func InstallGlobalFlagColorfulFlag(app *cli.Command) {
	app.Flags = append(app.Flags, colorfulFlag)
	before := app.Before
	app.Before = func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		if before != nil {
			_, err := before(ctx, cmd)
			if err != nil {
				return nil, err
			}
		}
		flag := cmd.Bool("colorful")
		if flag {
			colorful.O = colorful.Colorful{}
			printer.Printer = printer.NewWorker(printer.TypeColorful)
		}
		return nil, nil
	}
}

func InstallGlobalFlagJsonFlag(app *cli.Command) {
	app.Flags = append(app.Flags, jsonFlag)
	before := app.Before
	app.Before = func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		if before != nil {
			_, err := before(ctx, cmd)
			if err != nil {
				return nil, err
			}
		}
		flag := cmd.Bool("json")
		if flag {
			printer.Printer = printer.NewWorker(printer.TypeJson)
		}
		return nil, nil
	}
}

func InstallGlobalFlags(app *cli.Command) {
	InstallGlobalFlagDebug(app, log.Logger)
	InstallGlobalFlagExperimentalFlag(app)
	InstallGlobalFlagColorfulFlag(app)
	InstallGlobalFlagJsonFlag(app)
}
