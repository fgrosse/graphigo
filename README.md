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

```go
package main

import "github.com/fgrosse/graphigo"

func main() {
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
	metric := graphigo.Metric{Name: "test", Value: 3.14} // you can use any type as value
	defer client.Send(metric)

	// create a whole bunch of metrics and send them all with one TCP call
	client.SendAll([]graphigo.Metric{
		{Name: "shut", Value: 1},
		{Name: "up", Value: 2},
		{Name: "and", Value: 3},
		{Name: "take", Value: 4},
		{Name: "my", Value: 5},
		{Name: "money", Value: 6},
	})
}
```

**Note**: All exported functions of the graphigo client are noops if the client is `nil`. 

## Contributing

Any contributions are always welcome (use pull requests).
Please keep in mind that I might not always be able to respond immediately but I usually try to react within the week ☺.

[1]: http://graphite.readthedocs.org/en/latest/overview.html
[2]: http://godoc.org/github.com/fgrosse/graphigo
