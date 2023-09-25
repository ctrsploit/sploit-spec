package colorful

import (
	"fmt"
	"testing"
)

func Test_tickOrBallot(t *testing.T) {
	fmt.Printf("text: %s\n", Bool(Text{}, true))
	fmt.Printf("colorful: %s\n", Bool(Colorful{}, true))
}
