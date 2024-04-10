package app

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	env := Env{
		Basic{
			WebServer: WebServer{
				Name:    "nginx",
				Version: "1.0.0",
			},
			Framework: Framework{
				Name:    "spring",
				Version: "1.0.0",
			},
		},
		Advance{
			OS: OS{
				Type:    "linux",
				Version: "ubuntu 20.04",
			},
			ComponentList: []Component{
				{Name: "fastjson", Version: "1.0.0"},
				{Name: "test", Version: "1.1.0"},
			},
		},
	}
	marshaled, err := json.Marshal(env)
	assert.NoError(t, err)
	fmt.Printf("%s", marshaled)
}
