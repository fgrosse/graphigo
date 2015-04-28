Graphigo
========

This is simple go client for the [graphite monitoring tool][1].

## Installation

Use `go-get` to install graphite-golang
```
go get github.com/fgrosse/graphigo
```

No additional dependencies are required.

## Documentation

A generated documentation is available at [godoc.org][2]

## Usage

The following usage example is basically a copy of [tests/usage_example.go](tests/usage_example.go) for your convenience.

```go
package tests

import "github.com/fgrosse/graphigo"

func usageExample() {
    config = loadYourConfiguration()
    
    var client graphigo.GraphiteClient
    if config.GraphiteEnabled {
	    client = graphigo.NewClient("graphite.your.org:2003")
	    
	    // set a custom timeout for the graphite connection (optional, seconds, 0 = disabled)
        client.Timeout = 0
        
        // set a custom prefix for all recorded metrics of this client (optional)
        client.Prefix = "foo.bar.baz"
    } else {
        // there is also a null implementation which does not send any data to graphite
        client = graphigo.NewNullClient()
    }

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
```

## Contributing

Any contributions are always welcome (use pull requests).
Please keep in mind that I might not always be able to respond immediately but I usually try to react within the week ☺.

[1]: http://graphite.readthedocs.org/en/latest/overview.html
[2]: http://godoc.org/github.com/FGrosse/graphigo
