package graphigo_test

import "github.com/fgrosse/graphigo"

func Example() {
	client := graphigo.NewClient("graphite.your.org:2003")
	client.UseConnection(newConnectionMock())

	// set a custom timeout (seconds) for the graphite connection
	// if timeout = 0 then the graphigo.DefaultTimeout = 5 seconds is used
	// Setting Timeout to -1 disables the timeout
	client.Timeout = 0

	// set a custom prefix for all recorded metrics of this client (optional)
	client.Prefix = "foo.bar.baz"

	if err := client.Connect(); err != nil {
		panic(err)
	}

	// close the TCP connection properly if you don't need it anymore
	defer client.Disconnect()

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
}
