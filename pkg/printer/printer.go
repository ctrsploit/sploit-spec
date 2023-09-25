package printer

import "encoding/json"

const (
	TypeUnknown = iota
	TypeText
	TypeColorful
	TypeJson
	TypeDefault = TypeText
)

type Printer interface {
	Text() string
	Colorful() string
}

func Text(p Printer) (s string) {
	return p.Text()
}

func Colorful(p Printer) (s string) {
	return p.Colorful()
}

func Json(p Printer) (s string, err error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return
	}
	s = string(bytes)
	return
}
