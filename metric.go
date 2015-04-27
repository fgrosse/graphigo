package graphigo

import (
	"fmt"
	"time"
)

type Metric struct {
	Name      string
	Value     interface{}
	Timestamp time.Time
}

func NewMetric(name string, value interface{}) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: time.Now(),
	}
}

func (m Metric) UnixTimestamp() int64 {
	if m.Timestamp.IsZero() {
		m.Timestamp = time.Now()
	}

	return m.Timestamp.UTC().Unix()
}

func (m Metric) String() string {
	return fmt.Sprintf(
		"%s %v %s",
		m.Name,
		m.Value,
		m.Timestamp.UTC().Format("2006-01-02 15:04:05"),
	)
}
