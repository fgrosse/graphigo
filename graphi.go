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

// Client is a simple TCP client for the graphite monitoring tool.
type Client struct {
	// Address is used when connecting to the graphite server. Use address:port notation.
	Address string

	// Timeout is the maximum duration that the client will wait for a response from the server.
	// If timeout is zero then the DefaultTimeout is used.
	// Setting Timeout to -1 disables the timeout.
	Timeout time.Duration

	// Prefix is prepended to all metric names (separated from original name by a dot).
	// If Prefix is empty this does nothing.
	Prefix string

	// Connection is used to communicate with the graphite client.
	// You would normally not interfere with this but this can be handy for testing.
	Connection io.WriteCloser
}

const (
	// DefaultTimeout is the timeout that is applied when connecting to a graphite server.
	// It is used if no explicit timeout has been configured on a client.
	DefaultTimeout = 5

	// TimeoutDisabled is used to disable the client timeout entirely.
	TimeoutDisabled = -1
)

// NewClient creates a new instance of a graphite client.
// Use the address:port notation to specify the port.
// Note that the client will not connect to the given address automatically.
// You still need to call Connect() before you can start sending values.
func NewClient(address string) *Client {
	return &Client{Address: address}
}

// Connect attempts to establish the connection to the graphite server.
// This will return an error if a TCP connection can not or has already been established.
func (g *Client) Connect() (err error) {
	if g == nil {
		return nil
	}

	if g.Connection != nil {
		return fmt.Errorf("Graphigo is already connected. Call Disconnect first if you want to reconnect")
	}

	if g.Timeout == 0 {
		g.Timeout = DefaultTimeout * time.Second
	}

	timeout := g.Timeout
	if g.Timeout == -1 {
		timeout = 0
	}

	g.Connection, err = net.DialTimeout("tcp", g.Address, timeout)
	return err
}

// Disconnect closes the underlying connection to graphite.
func (g *Client) Disconnect() error {
	if g == nil {
		return nil
	}

	err := g.Connection.Close()
	g.Connection = nil
	return err
}

// SendValue creates a new graphigo.Metric with the metric timestamp set to now and sends it to graphite.
//
// Use Send(metric) if you want to split metric recording and sending.
// This will return an error if the client has not yet been connected or the metric name is empty.
func (g *Client) SendValue(name string, value interface{}) error {
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
func (g *Client) Send(metric Metric) error {
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
func (g *Client) SendAll(metrics []Metric) (err error) {
	if g == nil {
		return nil
	}

	if g.Connection == nil {
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

	_, err = g.Connection.Write(buffer.Bytes())
	return err
}
