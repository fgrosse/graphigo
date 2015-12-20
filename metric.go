package graphigo

import (
	"fmt"
	"time"
)

// Metric holds all information that is send to graphite.
// The value can be any go type but in practice graphite will probably only accept numeric values.
type Metric struct {
	// The Name of the metric.
	Name string

	// The Value of this metric. The client uses %v to format this.
	Value interface{}

	// Timestamp represents the time when this metric was recorded.
	// If this is the zero value the client will assume time.Now()
	Timestamp time.Time
}

// UnixTimestamp returns the the number of seconds elapsed since January 1, 1970 UTC.
// If the metrics timestamp is zero it will return time.Now().Unix()
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
