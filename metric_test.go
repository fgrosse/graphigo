package graphigo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"fmt"

	"gopkg.in/fgrosse/graphigo.v2"
)

var _ = Describe("Metric", func() {
	It("should implement fmt.Stringer", func() {
		var _ fmt.Stringer = graphigo.Metric{}
	})

	Describe("UnixTimestamp", func() {
		It("should use time.Now() if the TimeStamp is empty", func() {
			m := graphigo.Metric{}
			Expect(m.UnixTimestamp()).To(BeNumerically("~", time.Now().Unix(), 1))
		})
	})

	Describe("CaptureMetric", func() {
		It("should set the Timestamp to time.Now()", func() {
			m := graphigo.CaptureMetric("foo", 42)
			Expect(m.Timestamp).To(BeTemporally("~", time.Now(), time.Second))
		})
	})
})
