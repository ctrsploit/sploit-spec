package result

import "github.com/ctrsploit/sploit-spec/pkg/printer"

type Union struct {
	Machine printer.Interface
	Human   printer.Interface
}

func (p Union) Text() string {
	return p.Human.Text()
}

func (p Union) Colorful() string {
	return p.Human.Colorful()
}

func (p Union) Json() string {
	return printer.Json(p.Machine)
}

func (p Union) IsEmpty() bool {
	return false
}
