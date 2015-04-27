package graphigo

type GraphiteClient interface {
	Connect() error
	Disconnect() error
	Send(Metric) error
	SendAll([]Metric) error
}
