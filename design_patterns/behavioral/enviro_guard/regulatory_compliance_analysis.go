package main

import "fmt"

// ComplianceChecker is an interface representing a compliance checker in the chain.
type ComplianceChecker interface {
	CheckCompliance(data map[string]float64)
	SetNextChecker(nextChecker ComplianceChecker)
}

// BaseComplianceChecker represents the abstract class for compliance checkers.
type BaseComplianceChecker struct {
	nextChecker ComplianceChecker
}

func (bcc *BaseComplianceChecker) SetNextChecker(nextChecker ComplianceChecker) {
	bcc.nextChecker = nextChecker
}

// AirQualityComplianceChecker is a concrete implementation of a compliance checker for air quality.
type AirQualityComplianceChecker struct {
	BaseComplianceChecker
}

func (aqcc *AirQualityComplianceChecker) CheckCompliance(data map[string]float64) {
	if value, exists := data["Air Quality"]; exists && value > 40.0 {
		fmt.Println("Air quality exceeded threshold. Regulatory compliance violated.")
	} else if aqcc.nextChecker != nil {
		aqcc.nextChecker.CheckCompliance(data)
	}
}

// HumidityComplianceChecker is a concrete implementation of a compliance checker for humidity.
type HumidityComplianceChecker struct {
	BaseComplianceChecker
}

func (hcc *HumidityComplianceChecker) CheckCompliance(data map[string]float64) {
	if value, exists := data["Humidity"]; exists && value > 70.0 {
		fmt.Println("Humidity exceeded threshold. Regulatory compliance violated.")
	} else if hcc.nextChecker != nil {
		hcc.nextChecker.CheckCompliance(data)
	}
}

func main() {
	data := map[string]float64{
		"Temperature": 24.5,
		"Humidity":    73.2,
		"Air Quality": 38.7,
	}

	airQualityChecker := &AirQualityComplianceChecker{}
	humidityChecker := &HumidityComplianceChecker{}

	airQualityChecker.SetNextChecker(humidityChecker)

	airQualityChecker.CheckCompliance(data)
}
