Graphigo
========

[![Build Status](https://secure.travis-ci.org/fgrosse/graphigo.png?branch=master)](http://travis-ci.org/fgrosse/graphigo)
[![GoDoc](https://godoc.org/github.com/fgrosse/graphigo?status.svg)](https://godoc.org/github.com/fgrosse/graphigo)

A simple go client for the [graphite monitoring tool][1].

## Installation

Use `go get` to install graphigo:
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
    client := graphigo.NewClient("graphite.your.org:2003")

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
```

There is also the `NullClient` implementation which implements the `GraphiteClient` interface but does not actually send any data to graphite.

```go
package tests

import "github.com/fgrosse/graphigo"

// Note that we do return the interface so we can also set up the concrete implementation in this function
func setupGraphiteClient(address string, enabled bool) graphigo.GraphiteClient {
    if enabled {
	    client := graphigo.NewClient(address)
        client.Prefix = "foo.bar.baz"
        return client
    } else {
        return graphigo.NewNullClient()
    }
}
```

## Running the tests

Graphigo uses the awesome [ginkgo][3] framework for its tests.
You can execute the tests running:
```
ginkgo tests
```

If you prefer to use `go test` directly you can either switch into the `./tests` directory and run it there or
run the following from the repository root directory:
```
go test ./tests
```

## Contributing

Any contributions are always welcome (use pull requests).
Please keep in mind that I might not always be able to respond immediately but I usually try to react within the week ☺.

[1]: http://graphite.readthedocs.org/en/latest/overview.html
[2]: http://godoc.org/github.com/fgrosse/graphigo
[3]: http://onsi.github.io/ginkgo/
