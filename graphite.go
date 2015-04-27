package graphigo

// GraphiteClient is the interface for a graphigo graphite client.
// Use it for dependency injection and the good of humanity
type GraphiteClient interface {
	Connect() error
	Disconnect() error
	Send(Metric) error
	SendAll([]Metric) error
}
