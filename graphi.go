// Package graphigo provides a simple go client for the graphite monitoring tool.
// See http://graphite.readthedocs.org/en/latest/overview.html
package graphigo

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

// Graphigo is a simple implementation of the GraphiteClient interface
// You can get new instances using graphigo.NewClient(address)
//
// If you want to configure a global name prefix for all recorded values you
// can access the globally available Prefix field.
//
// To set a timeout on the underlying connection to graphite, set the Timeout field
// of this struct.
type Graphigo struct {
	Address string
	Timeout time.Duration
	Prefix  string

	connection io.WriteCloser
}

const (
	// DefaultTimeout is the timeout that is applied when connecting to a graphite server
	// if no explicit timeout has been configured.
	DefaultTimeout = 5

	// TimeoutDisabled is used to disable the client timeout entirely
	TimeoutDisabled = -1
)

// NewClient creates a new instance of a graphite client.
// Use the address:port notation to specify the port.
func NewClient(address string) *Graphigo {
	return &Graphigo{Address: address}
}

// UseConnection can be used to inject your own implementation of the connection to graphite.
// The connection is closed by any call to Disconnect().
func (g *Graphigo) UseConnection(c io.WriteCloser) {
	g.connection = c
}

// Connect attempts to establish the connection to the graphite server.
// This will return an error if a TCP connection can not or has already been established.
func (g *Graphigo) Connect() (err error) {
	if g == nil {
		return nil
	}

	if g.connection != nil {
		return fmt.Errorf("Graphigo is already connected. Call Disconnect first if you want to reconnect")
	}

	if g.Timeout == 0 {
		g.Timeout = DefaultTimeout * time.Second
	}

	timeout := g.Timeout
	if g.Timeout == -1 {
		timeout = 0
	}

	g.connection, err = net.DialTimeout("tcp", g.Address, timeout)
	return err
}

// Disconnect closes the underlying connection to graphite.
func (g *Graphigo) Disconnect() error {
	if g == nil {
		return nil
	}

	err := g.connection.Close()
	g.connection = nil
	return err
}

// SendValue creates a new graphigo.Metric with the metric timestamp set to now and sends it to graphite.
//
// Use Send(metric) if you want to split metric recording and sending.
// This will return an error if the client has not yet been connected or the metric name is empty.
func (g *Graphigo) SendValue(name string, value interface{}) error {
	if g == nil {
		return nil
	}

	return g.Send(Metric{
		Name:  name,
		Value: value,
	})
}

// Send sends a graphigo.Metric to graphite.
// This can be used to send a metric which has been recorded earlier.
//
// Use SendValue if you want to create and send a metric in one step.
// Use SendAll if you want to send multiple metrics at once.
// This will return an error if the client has not yet been connected or the metric name is empty.
func (g *Graphigo) Send(metric Metric) error {
	if g == nil {
		return nil
	}

	return g.SendAll([]Metric{metric})
}

// SendAll sends multiple graphigo.Metric to graphite.
// This can be used to send multiple metrics that have been recorded earlier.
//
// Use Send if you want to send a single metric.
// This will return an error if the client has not yet been connected or if any of the metrics has an empty name.
func (g *Graphigo) SendAll(metrics []Metric) (err error) {
	if g == nil {
		return nil
	}

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

		_, err = fmt.Fprintf(buffer, "%s %v %d\n", metric.Name, metric.Value, metric.UnixTimestamp())
		if err != nil {
			return fmt.Errorf("Could not write metric to send buffer: %s", err.Error())
		}
	}

	_, err = g.connection.Write(buffer.Bytes())
	return err
}
