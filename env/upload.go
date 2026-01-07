package env

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ctrsploit/sploit-spec/pkg/env/linux"
	"github.com/ctrsploit/sploit-spec/pkg/upload"
	"github.com/urfave/cli/v3"
)

var LinuxEnv linux.Env

func UploadAction(context context.Context, cmd *cli.Command) (err error) {
	//eg. ECS
	servicename := cmd.Args().Get(0)
	// region_tag.json eg. cn-north4_linux.json
	filename := cmd.Args().Get(1)
	// obsurl
	obsurl := cmd.Args().Get(2)
	// obshost (if want to hide obs upload behavior), put your real obsurl in here, put the fake url in obsurl
	obshost := cmd.Args().Get(3)
	if servicename == "" {
		return
	}

	filename = servicename + "_" + filename

	// write collect your env info code in here
	// ...

	// compiletime, err := kernel.GetKernelCompileTime()
	// if err != nil{
	// 	fmt.Println("Failed to Get Kernel Compile Time")
	// }
	// compiletimestr := compiletime.Format("2006-01-02 15:04:05")
	// LinuxEnv.Basic.CompileTime = compiletimestr

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	// Need write the Timestamp and ServiceName to your info struct
	LinuxEnv.TimeStamp = timeString
	LinuxEnv.ServiceName = servicename

	// Convert your env struct to json bytes
	resstring, err := json.Marshal(LinuxEnv)
	if err != nil {
		fmt.Println("Json.Marshal failed", err)
		return err
	}
	// Finally call uploadtobs
	err = upload.UploadToHostObs(resstring, obsurl, obshost, filename)
	if err != nil {
		fmt.Println("Upload to Obs failed")
	}

	return
}

var Upload = &cli.Command{
	Name:    "upload",
	Aliases: []string{"up"},
	Usage:   "upload env servicename filename obsurl obshost",
	Action:  UploadAction,
}
