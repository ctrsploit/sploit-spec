package item

type Item interface {
	Text() string
	Colorful() string
}

func Text(i Item) (s string) {
	return i.Text()
}

func Colorful(i Item) (s string) {
	return i.Colorful()
}
