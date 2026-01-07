package upload

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func GenerateUploadCommand(env func() (json []byte, err error)) (cmd *cli.Command) {
	return &cli.Command{
		Name:    "upload",
		Aliases: []string{"up"},
		Usage:   "upload <servicename> <filename> <obs> [host]",

		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
			if cmd.NArg() < 3 {
				return cli.Exit(fmt.Errorf("invalid arguments"), 1)
			}
			//eg. ECS
			servicename := cmd.Args().Get(0)
			// region_tag.json eg. cn-north4_linux.json
			filename := cmd.Args().Get(1)
			// obsurl
			obs := cmd.Args().Get(2)
			// obshost (if want to hide obs upload behavior), put your real obsurl in here, put the fake url in obsurl
			host := cmd.Args().Get(3)
			if servicename == "" {
				return
			}
			filename = servicename + "_" + filename
			json, err := env()
			if err != nil {
				return
			}
			err = Obs(json, filename, obs, host)
			if err != nil {
				fmt.Println("Upload to Obs failed")
				return
			}
			return
		},
	}
}
