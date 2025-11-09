package env

import (
	"encoding/json"
	"xsploit/env/auto"

	"github.com/ctrsploit/sploit-spec/pkg/upload"
)

var (
	Upload = upload.GenerateUploadCommand(func() (content []byte, err error) {
		env := auto.Auto()
		content, err = json.Marshal(env)
		if err != nil {
			return
		}
		return
	})
)
