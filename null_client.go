package graphigo

// NullClient implements the GraphiteClient but does not perform any actions
type NullClient struct{}

func NewNullClient() *NullClient {
	return &NullClient{}
}

func (c *NullClient) Connect() error {
	return nil
}

func (c *NullClient) Disconnect() error {
	return nil
}

func (c *NullClient) SendValue(_ string, _ interface{}) error {
	return nil
}

func (c *NullClient) Send(Metric) error {
	return nil
}

func (c *NullClient) SendAll([]Metric) error {
	return nil
}
