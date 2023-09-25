package result

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
