package main

import (
	"fmt"
)

// HistoricalDataAnalysisStrategy defines the interface for historical data analysis strategies.
type HistoricalDataAnalysisStrategy interface {
	Analyze(data map[string][]float64) map[string]interface{}
}

// CalculateAverageStrategy is a concrete strategy to calculate the average of sensor data.
type CalculateAverageStrategy struct{}

func (cas *CalculateAverageStrategy) Analyze(data map[string][]float64) map[string]interface{} {
	results := make(map[string]interface{})
	for sensor, values := range data {
		sum := 0.0
		for _, value := range values {
			sum += value
		}
		average := sum / float64(len(values))
		results[sensor] = average
	}
	return results
}

// IdentifyPatternsStrategy is a concrete strategy to identify patterns in sensor data.
type IdentifyPatternsStrategy struct{}

func (ips *IdentifyPatternsStrategy) Analyze(data map[string][]float64) map[string]interface{} {
	results := make(map[string]interface{})
	const threshold = 100.0
	for sensor, values := range data {
		pattern := "Stable"
		for _, value := range values {
			if value >= threshold {
				pattern = "Variable"
				break
			}
		}
		results[sensor] = pattern
	}
	return results
}

// DataAnalyzer is a context class that uses a historical data analysis strategy.
type DataAnalyzer struct {
	strategy HistoricalDataAnalysisStrategy
}

func (da *DataAnalyzer) SetStrategy(strategy HistoricalDataAnalysisStrategy) {
	da.strategy = strategy
}

func (da *DataAnalyzer) AnalyzeData(data map[string][]float64) map[string]interface{} {
	return da.strategy.Analyze(data)
}

func main() {
	data := map[string][]float64{
		"Air Quality":  {90, 110, 85, 95, 120},
		"Noise Level":  {70, 75, 80, 85, 90},
	}

	averageStrategy := &CalculateAverageStrategy{}
	patternStrategy := &IdentifyPatternsStrategy{}

	analyzer := &DataAnalyzer{strategy: averageStrategy}
	averageResults := analyzer.AnalyzeData(data)
	fmt.Println("Average Results:", averageResults)

	analyzer.SetStrategy(patternStrategy)
	patternResults := analyzer.AnalyzeData(data)
	fmt.Println("Pattern Results:", patternResults)
}
