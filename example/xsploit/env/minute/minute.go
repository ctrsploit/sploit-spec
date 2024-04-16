package minute

import (
	"time"
)

type Spec struct {
	Minute int `json:"minute"`
}

func Minute() (minute Spec) {
	minute = Spec{
		Minute: time.Now().Minute(),
	}
	return
}
