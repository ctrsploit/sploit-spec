package app

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_error/exporter"
	libraryLogger "github.com/ssst0n3/awesome_libs/log"
	"github.com/urfave/cli/v2"
	"io"
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
)

func InstallGlobalFlagDebug(app *cli.App, appLogger *logrus.Logger) {
	app.Flags = append(app.Flags, debugFlag)
	before := app.Before
	app.Before = func(context *cli.Context) (err error) {
		if before != nil {
			err = before(context)
			if err != nil {
				return
			}
		}
		debug := context.Bool("debug")
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
			libraryLogger.Logger.Level = logrus.DebugLevel
			libraryLogger.Logger.Debug("debug mode on")
		}
		return
	}
}

func InstallGlobalFlagExperimentalFlag(app *cli.App) {
	app.Flags = append(app.Flags, experimentalFlag)
}

func InstallGlobalFlagColorfulFlag(app *cli.App) {
	app.Flags = append(app.Flags, colorfulFlag)
	before := app.Before
	app.Before = func(ctx *cli.Context) (err error) {
		if before != nil {
			err = before(ctx)
			if err != nil {
				return
			}
		}
		flag := ctx.Bool("colorful")
		if flag {
			colorful.O = colorful.Colorful{}
			Printer = printer.GetPrinter(printer.TypeColorful)
		}
		return
	}
}

func InstallGlobalFlags(app *cli.App) {
	InstallGlobalFlagDebug(app, log.Logger)
	InstallGlobalFlagExperimentalFlag(app)
	InstallGlobalFlagColorfulFlag(app)
}
