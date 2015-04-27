package graphigo

import (
	"fmt"
	"time"
)

// Metric holds all information that is send to graphite.
// The value can be any go type but in practice graphite will probably only accept numeric values.
type Metric struct {
	Name      string
	Value     interface{}
	Timestamp time.Time
}

// NewMetric creates a new metric
func NewMetric(name string, value interface{}) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: time.Now(),
	}
}

// UnixTimestamp returns the the number of seconds elapsed since January 1, 1970 UTC.
func (m Metric) UnixTimestamp() int64 {
	if m.Timestamp.IsZero() {
		m.Timestamp = time.Now()
	}

	return m.Timestamp.Unix()
}

// String returns a textual representation of a metric (used for debugging)
func (m Metric) String() string {
	return fmt.Sprintf(
		"%s %v %s",
		m.Name,
		m.Value,
		m.Timestamp.UTC().Format("2006-01-02 15:04:05"),
	)
}
