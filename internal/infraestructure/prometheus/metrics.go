package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/zuluapp/go-libs/pkg/utils"
)

type (
	Metrics struct {
		registry prometheus.Registerer
		factory  promauto.Factory
		names    metricsNames
	}

	metricsNames map[string]prometheus.Metric

	metricType string
)

const (
	counter metricType = "counter"
	timer   metricType = "timer"
	gauge   metricType = "gauge"
)

func InitMetrics() Metrics {
	factory := promauto.With(prometheus.DefaultRegisterer)

	return Metrics{
		registry: prometheus.DefaultRegisterer,
		factory:  factory,
		names:    map[string]prometheus.Metric{},
	}
}

func (m Metrics) Increment(name string, tags utils.Map) {
	counterMetric := m.getMetric(name, tags, counter).(prometheus.Counter)
	counterMetric.Inc()
}

func (m Metrics) ObserveDuration(name string, tags utils.Map) func() time.Duration {
	histogramMetric := m.getMetric(name, tags, timer).(prometheus.Histogram)
	timerMetric := prometheus.NewTimer(histogramMetric)

	return func() time.Duration {
		return timerMetric.ObserveDuration()
	}
}

func (m Metrics) getMetric(name string, tags utils.Map, mtype metricType) prometheus.Metric {
	if metric, exists := m.names[name]; exists {
		return metric
	}

	return m.createMetric(name, tags, mtype)
}

func (m Metrics) createMetric(name string, tags utils.Map, mtype metricType) prometheus.Metric {
	switch mtype {
	case counter:
		counterMetric := m.factory.NewCounter(prometheus.CounterOpts{
			Name:        name,
			ConstLabels: createLabels(tags),
			Namespace:   "prometheus_go_app",
		})
		m.names[name] = counterMetric
		return counterMetric
	case timer:
		timerMetric := m.factory.NewHistogram(prometheus.HistogramOpts{
			Name:        name,
			ConstLabels: createLabels(tags),
			Namespace:   "prometheus_go_app",
		})
		m.names[name] = timerMetric
		return timerMetric
	}

	return nil
}

func createLabels(tags utils.Map) prometheus.Labels {
	labels := prometheus.Labels{}

	for key, value := range tags {
		switch val := value.(type) {
		case string:
			labels[key] = val
		}
	}

	return labels
}
