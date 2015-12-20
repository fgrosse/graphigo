package graphigo_test

import "gopkg.in/fgrosse/graphigo.v2"

func Example() {
	c := graphigo.NewClient("graphite.your.org:2003")
	c.Connection = newConnectionMock()

	// set a custom timeout (seconds) for the graphite connection
	// if timeout = 0 then the graphigo.DefaultTimeout = 5 seconds is used
	// Setting Timeout to -1 disables the timeout
	c.Timeout = 0

	// set a custom prefix for all recorded metrics of this client (optional)
	c.Prefix = "foo.bar.baz"

	if err := c.Connect(); err != nil {
		panic(err)
	}

	// close the TCP connection properly if you don't need it anymore
	defer c.Close()

	// capture and send values using a single line
	c.SendValue("hello.graphite.world", 42)

	// capture a metric and send it any time later
	metric := graphigo.Metric{Name: "test", Value: 3.14} // you can use any type as value
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
