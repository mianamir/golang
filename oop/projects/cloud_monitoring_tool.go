package main

import "fmt"

// MetricProvider is the interface defining the contract for fetching metric data.
type MetricProvider interface {
	GetMetricData(metricName string) float64
}

// AWSMetricProvider represents a metric provider for AWS CloudWatch.
type AWSMetricProvider struct{}

// GetMetricData simulates fetching metric data from AWS CloudWatch.
func (a *AWSMetricProvider) GetMetricData(metricName string) float64 {
	return 100.0
}

// AzureMetricProvider represents a metric provider for Azure Monitor.
type AzureMetricProvider struct{}

// GetMetricData simulates fetching metric data from Azure Monitor.
func (a *AzureMetricProvider) GetMetricData(metricName string) float64 {
	return 150.0
}

// MetricAggregator is a singleton that aggregates metric data.
type MetricAggregator struct {
	metrics map[string]float64
}

var metricAggregatorInstance *MetricAggregator

// GetMetricAggregatorInstance returns the singleton instance of MetricAggregator.
func GetMetricAggregatorInstance() *MetricAggregator {
	if metricAggregatorInstance == nil {
		metricAggregatorInstance = &MetricAggregator{
			metrics: make(map[string]float64),
		}
	}
	return metricAggregatorInstance
}

// AddMetric adds a metric to the aggregator.
func (ma *MetricAggregator) AddMetric(metricName string, metricValue float64) {
	ma.metrics[metricName] = metricValue
}

// GetMetrics returns aggregated metrics.
func (ma *MetricAggregator) GetMetrics() map[string]float64 {
	return ma.metrics
}

// CloudMonitoringSystem represents the cloud monitoring system.
type CloudMonitoringSystem struct {
	metricProviders  []MetricProvider
	metricAggregator *MetricAggregator
}

// NewCloudMonitoringSystem creates a new CloudMonitoringSystem instance.
func NewCloudMonitoringSystem(metricProviders []MetricProvider, metricAggregator *MetricAggregator) *CloudMonitoringSystem {
	return &CloudMonitoringSystem{
		metricProviders:  metricProviders,
		metricAggregator: metricAggregator,
	}
}

// FetchAndAggregateMetrics fetches and aggregates metrics.
func (cms *CloudMonitoringSystem) FetchAndAggregateMetrics(metricNames []string) {
	for _, provider := range cms.metricProviders {
		for _, metricName := range metricNames {
			metricValue := provider.GetMetricData(metricName)
			cms.metricAggregator.AddMetric(fmt.Sprintf("%T_%s", provider, metricName), metricValue)
		}
	}
}

// DisplayMetrics displays aggregated metrics.
func (cms *CloudMonitoringSystem) DisplayMetrics() {
	metrics := cms.metricAggregator.GetMetrics()
	for metricName, metricValue := range metrics {
		fmt.Printf("Metric: %s, Value: %f\n", metricName, metricValue)
	}
}

func main() {
	// Creating metric providers
	awsMetricProvider := &AWSMetricProvider{}
	azureMetricProvider := &AzureMetricProvider{}

	// Creating metric aggregator
	metricAggregator := GetMetricAggregatorInstance()

	// Creating cloud monitoring system
	cloudMonitoringSystem := NewCloudMonitoringSystem([]MetricProvider{awsMetricProvider, azureMetricProvider}, metricAggregator)

	// Fetching and aggregating metrics
	metricNames := []string{"CPUUsage", "MemoryUsage"}
	cloudMonitoringSystem.FetchAndAggregateMetrics(metricNames)

	// Displaying aggregated metrics
	cloudMonitoringSystem.DisplayMetrics()
}
