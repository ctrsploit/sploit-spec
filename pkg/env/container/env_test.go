package container

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvJson(t *testing.T) {
	env := Env{}
	marshaled, err := json.Marshal(env)
	assert.NoError(t, err)
	fmt.Printf("%s", marshaled)
}
