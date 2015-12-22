package graphigo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"gopkg.in/fgrosse/graphigo.v2"
)

var _ = Describe("Graphigo", func() {
	var (
		c    *graphigo.Client
		conn *connMock
	)

	BeforeEach(func() {
		conn = newConnectionMock()
		c = &graphigo.Client{}
		c.Connection = conn
	})

	Describe("Close", func() {
		It("should close the connection", func() {
			Expect(conn.IsClosed).NotTo(BeTrue())
			c.Close()
			Expect(conn.IsClosed).To(BeTrue())
		})
	})

	Describe("Send", func() {
		It("should send string values to graphite", func() {
			Expect(c.Send(graphigo.Metric{Name: "test_metric", Value: "42"})).To(Succeed())
			Expect(conn.SentMetrics).To(HaveLen(1))
			Expect(conn.SentMetrics[0].Name).To(Equal("test_metric"))
			Expect(conn.SentMetrics[0].Value).To(Equal("42"))
			Expect(conn.SentMetrics[0].Timestamp).To(BeTemporally("~", time.Now().UTC(), 1*time.Second))
		})

		It("should send integer values to graphite", func() {
			Expect(c.Send(graphigo.Metric{Name: "test_metric", Value: 7})).To(Succeed())
			Expect(conn.SentMetrics).To(HaveLen(1))
			Expect(conn.SentMetrics[0].Name).To(Equal("test_metric"))
			Expect(conn.SentMetrics[0].Value).To(Equal("7"))
			Expect(conn.SentMetrics[0].Timestamp).To(BeTemporally("~", time.Now().UTC(), 1*time.Second))
		})

		It("should send float values to graphite", func() {
			Expect(c.Send(graphigo.Metric{Name: "test_metric", Value: 3.14159265359})).To(Succeed())
			Expect(conn.SentMetrics).To(HaveLen(1))
			Expect(conn.SentMetrics[0].Name).To(Equal("test_metric"))
			Expect(conn.SentMetrics[0].Value).To(Equal("3.14159265359"))
			Expect(conn.SentMetrics[0].Timestamp).To(BeTemporally("~", time.Now().UTC(), 1*time.Second))
		})

		Context("with prefix", func() {
			BeforeEach(func() {
				c.Prefix = "foo_bar.baz"
			})

			It("should set the correct metric name", func() {
				Expect(c.Send(graphigo.Metric{Name: "test_metric", Value: 7})).To(Succeed())
				Expect(conn.SentMetrics).To(HaveLen(1))
				Expect(conn.SentMetrics[0].Name).To(Equal("foo_bar.baz.test_metric"))
			})

			It("should not leak changes to the given metric", func() {
				metric := graphigo.Metric{Name: "test_metric", Value: 7}
				Expect(c.Send(metric)).To(Succeed())
				Expect(metric.Name).To(Equal("test_metric"))
			})
		})
	})

	Describe("SendAll", func() {
		It("should send multiple values in one write command", func() {
			metrics := []graphigo.Metric{
				{Name: "test_metric_a", Value: "1", Timestamp: time.Now().Add(-1 * time.Hour)},
				{Name: "test_metric_b", Value: "2", Timestamp: time.Now().Add(-30 * time.Minute)},
				{Name: "test_metric_c", Value: "3", Timestamp: time.Now()},
			}
			Expect(c.SendAll(metrics)).To(Succeed())

			for i, sentMetric := range conn.SentMetrics {
				Expect(sentMetric.Name).To(Equal(metrics[i].Name))
				Expect(sentMetric.Value).To(Equal(metrics[i].Value))
				Expect(sentMetric.Timestamp).To(BeTemporally("~", metrics[i].Timestamp, 1*time.Second))
			}
		})
	})
})
