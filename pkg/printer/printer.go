package printer

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

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
	IsEmpty() bool
}

func Text(p Printer) (s string) {
	return p.Text()
}

func Colorful(p Printer) (s string) {
	return p.Colorful()
}

func Json(p Printer) (s string) {
	bytes, err := json.Marshal(p)
	if err != nil {
		awesome_error.CheckWarning(err)
		return
	}
	s = string(bytes)
	return
}

func GetPrinter(t int) func(p Printer) (s string) {
	switch t {
	case TypeText:
		return Text
	case TypeColorful:
		return Colorful
	case TypeJson:
		return Json
	}
	return nil
}
