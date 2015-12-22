package graphigo_test

import (
	"gopkg.in/fgrosse/graphigo.v2"
	"time"
)

func Example() {
	c := graphigo.Client{
		// If you omit the entire address localhost:2004 will be assumed
		// Just omitting the port is also valid and wil use the default port
		Address: "graphite.your.org:2004",

		// set a custom timeout (seconds) for the graphite connection
		// if timeout = 0 then the graphigo.DefaultTimeout = 5 seconds is used
		// Setting Timeout to graphite.TimeoutDisabled (-1) disables the timeout
		Timeout: 42,

		// set a custom prefix for all recorded metrics of this client (optional)
		Prefix: "foo.bar.baz",
	}

	c.Connection = newConnectionMock()
	if err := c.Connect(); err != nil {
		panic(err) // do proper error handling
	}

	// close the TCP connection properly if you don't need it anymore
	defer c.Close()

	// capture and send values using a single line
	c.SendValue("hello.graphite.world", 42)

	// capture a metric and send it any time later. You can use any type as value
	metric := graphigo.Metric{Name: "test", Value: 3.14, Timestamp: time.Now()}
	defer c.Send(metric)

	// create a whole bunch of metrics and send them all with one TCP call
	c.SendAll([]graphigo.Metric{
		{Name: "shut", Value: 1},
		{Name: "up", Value: 2},
		{Name: "and", Value: 3},
		{Name: "take", Value: 4},
		{Name: "my", Value: 5},
		{Name: "money", Value: 6},
	})
}
