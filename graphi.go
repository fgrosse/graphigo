package graphigo

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

type Graphigo struct {
	Address string
	Timeout time.Duration
	Prefix  string

	connection io.WriteCloser
	nop        bool
}

const DefaultTimeout = 5

func New(address string) *Graphigo {
	return &Graphigo{Address: address}
}

func (g *Graphigo) InjectConnection(connection io.WriteCloser) {
	g.connection = connection
}

func (g *Graphigo) Connect() (err error) {
	if g.connection != nil {
		return fmt.Errorf("Graphigo is already connected. Call Disconnect first if you want to reconnect")
	}

	if g.Timeout == 0 {
		g.Timeout = DefaultTimeout * time.Second
	}

	g.connection, err = net.DialTimeout("tcp", g.Address, g.Timeout)
	return err
}

func (g *Graphigo) Disconnect() error {
	err := g.connection.Close()
	g.connection = nil
	return err
}

func (g *Graphigo) SendValue(name string, value interface{}) error {
	return g.Send(Metric{
		Name:  name,
		Value: value,
	})
}

func (g *Graphigo) Send(metric Metric) error {
	return g.SendAll([]Metric{metric})
}

func (g *Graphigo) SendAll(metrics []Metric) error {
	if g.connection == nil {
		return fmt.Errorf("Graphigo is not connected yet. Did you forget to call Connect() ?")
	}

	buffer := &bytes.Buffer{}
	for _, metric := range metrics {
		if metric.Name == "" {
			return fmt.Errorf("Could not send graphite metric: no metric name given")
		}

		if metric.Timestamp.IsZero() {
			metric.Timestamp = time.Now()
		}

		if g.Prefix != "" {
			metric.Name = fmt.Sprintf("%s.%s", g.Prefix, metric.Name)
		}

		fmt.Fprintf(buffer, "%s %v %d\n", metric.Name, metric.Value, metric.UnixTimestamp())
	}

	_, err := g.connection.Write(buffer.Bytes())
	return err
}
