package solid_principles

type Orderable interface {
	GenerateInvoice(total float64) string
}

type OrderProcessor struct{}

func (op *OrderProcessor) ProcessOrder(orderable Orderable, total float64) string {
	invoice := orderable.GenerateInvoice(total)
	return invoice
}
