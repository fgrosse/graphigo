package tests

import (
	. "github.com/onsi/ginkgo"

	"github.com/fgrosse/graphigo"
)

var _ = Describe("Usage Example", func() {

	It("should not panic", func() {
		conn := newConnectionMock()
		client := &graphigo.Graphigo{}
		client.InjectConnection(conn)

		// set a custom timeout for the graphite connection (optional, seconds, 0 = disabled)
		client.Timeout = 0

		if false { // calling this in the test were we already injected a connection will (correctly) cause an error)
			if err := client.Connect(); err != nil {
				panic(err)
			}
		}

		// close the TCP connection properly if you don't need it anymore
		defer client.Disconnect()

		// set a custom prefix for all recorded metrics of this client (optional)
		client.Prefix = "foo.bar.baz"

		// capture and send values using a single line
		client.SendValue("hello.graphite.world", 42)

		// capture a metric and send it any time later
		metric := graphigo.NewMetric("test", 3.14) // you can use any type as value
		defer client.Send(metric)

		// create a whole bunch of metrics and send them all with one TCP call
		metrics := []graphigo.Metric{
			graphigo.NewMetric("foo", 1),
			graphigo.NewMetric("bar", 1.23),
			graphigo.NewMetric("baz", "456"),
		}
		client.SendAll(metrics)

		// of course this all works in once line and still reads nicely
		client.SendAll([]graphigo.Metric{
			graphigo.NewMetric("shut", 1),
			graphigo.NewMetric("up", 2),
			graphigo.NewMetric("and", 3),
			graphigo.NewMetric("take", 4),
			graphigo.NewMetric("my", 5),
			graphigo.NewMetric("money", 6),
		})
	})

	It("should work with the NullClient", func() {
		client := graphigo.NewNullClient()

		if err := client.Connect(); err != nil {
			panic(err)
		}

		defer client.Disconnect()

		client.SendValue("hello.graphite.world", 42)
		metric := graphigo.NewMetric("test", 3.14) // you can use any type as value
		client.Send(metric)

		metrics := []graphigo.Metric{
			graphigo.NewMetric("foo", 1),
			graphigo.NewMetric("bar", 1.23),
			graphigo.NewMetric("baz", "456"),
		}
		client.SendAll(metrics)
	})
})
