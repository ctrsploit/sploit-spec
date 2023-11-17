package printer

type Worker struct {
	Type      int
	PrintFunc PrintFunc
}

func NewWorker(t int) *Worker {
	return &Worker{
		Type:      t,
		PrintFunc: GetPrinter(t),
	}
}

func (w Worker) Print(printers ...Printer) (s string) {
	s = Print(w.PrintFunc, printers...)
	return
}
