package grasp_patterns

import (
	"fmt"
	"time"
)

// MetricType represents the type of metric being monitored
type MetricType string

// Metric represents a metric data point in the monitoring system
type Metric struct {
	Timestamp time.Time
	Type      MetricType
	Value     float64
}

// MetricStore is responsible for storing and managing metric data
type MetricStore struct {
	metrics map[MetricType][]Metric
}

func NewMetricStore() *MetricStore {
	return &MetricStore{
		metrics: make(map[MetricType][]Metric),
	}
}

func (ms *MetricStore) AddMetric(metric Metric) {
	ms.metrics[metric.Type] = append(ms.metrics[metric.Type], metric)
	fmt.Printf("Added metric - Type: %s, Value: %.2f\n", metric.Type, metric.Value)
}

func main() {
	metricStore := NewMetricStore()

	metric1 := Metric{Timestamp: time.Now(), Type: "CPU", Value: 0.85}
	metric2 := Metric{Timestamp: time.Now(), Type: "Memory", Value: 0.72}

	metricStore.AddMetric(metric1)
	metricStore.AddMetric(metric2)
}
