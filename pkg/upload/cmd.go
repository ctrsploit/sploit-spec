package upload

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func GenerateUploadCommand(env func() (json []byte, err error)) (cmd *cli.Command) {
	return &cli.Command{
		Name:    "upload",
		Aliases: []string{"up"},
		Usage:   "upload <servicename> <filename> <obsurl> <obshost>",
		Action: func(context *cli.Context) (err error) {
			//eg. ECS
			servicename := context.Args().Get(0)
			// region_tag.json eg. cn-north4_linux.json
			filename := context.Args().Get(1)
			// obsurl
			obs := context.Args().Get(2)
			// obshost (if want to hide obs upload behavior), put your real obsurl in here, put the fake url in obsurl
			host := context.Args().Get(3)
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