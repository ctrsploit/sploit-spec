package second

import (
	"time"
)

type Spec struct {
	Second int `json:"second"`
}

func Second() (second Spec) {
	second = Spec{
		Second: time.Now().Second(),
	}
	return
}
