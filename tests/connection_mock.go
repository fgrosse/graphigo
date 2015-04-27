package tests

import (
	"fmt"
	"github.com/fgrosse/graphigo"
	"strconv"
	"strings"
	"time"
)

type connectionMock struct {
	IsClosed    bool
	SentMetrics []graphigo.Metric
}

func newConnectionMock() *connectionMock {
	return &connectionMock{
		IsClosed:    false,
		SentMetrics: []graphigo.Metric{},
	}
}

func (c *connectionMock) Write(b []byte) (n int, err error) {
	metricLines := strings.Split(string(b), "\n")
	if len(metricLines) == 0 {
		return 0, fmt.Errorf("No metrics given at all!")
	}

	for i, line := range metricLines {
		if line == "" {
			continue
		}

		metricParts := strings.Split(line, " ")
		if len(metricParts) != 3 {
			return 0, fmt.Errorf("Invalid metric format given in metric line %d: %q", i, line)
		}

		newMetric := graphigo.Metric{Name: metricParts[0], Value: metricParts[1]}
		timeStamp, err := strconv.ParseInt(metricParts[2], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("Could not parse metric timestamp %q: %s", metricParts[2], err.Error())
		}

		newMetric.Timestamp = time.Unix(timeStamp, 0).UTC()
		c.SentMetrics = append(c.SentMetrics, newMetric)
	}

	return len(b), nil
}

func (c *connectionMock) Close() error {
	c.IsClosed = true
	return nil
}
