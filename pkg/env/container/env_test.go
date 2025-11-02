package container

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestEnvJson(t *testing.T) {
	env := Env{}
	spew.Dump(env)
	marshaled, err := json.Marshal(env)
	assert.NoError(t, err)
	fmt.Printf("%s", marshaled)
}
