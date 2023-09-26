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

type PrintFunc func(p interface{}) (s string)

func Text(p interface{}) (s string) {
	return p.(Printer).Text()
}

func Colorful(p interface{}) (s string) {
	return p.(Printer).Colorful()
}

func Json(p interface{}) (s string) {
	bytes, err := json.Marshal(p)
	if err != nil {
		awesome_error.CheckWarning(err)
		return
	}
	s = string(bytes)
	return
}

func GetPrinter(t int) PrintFunc {
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

func Print(print PrintFunc, printers ...Printer) (s string) {
	for _, i := range printers {
		if !i.IsEmpty() {
			s += print(i) + "\n"
		}
	}
	return
}
