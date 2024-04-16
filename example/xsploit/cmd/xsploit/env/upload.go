package env

import (
	"encoding/json"
	"github.com/ctrsploit/sploit-spec/pkg/upload"
	"xsploit/env/auto"
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
