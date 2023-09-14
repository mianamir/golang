package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Define a set of allowed functions and operations
var allowedFunctions = map[string]bool{
	"filter": true,
	"map":    true,
	"sum":    true,
	"min":    true,
	"max":    true,
	"lambda": true,
	"data":   true,
}

func loadUserData() [][]int {
	// Load user data (sample data for demonstration).
	// Modify this function to load actual data sources.
	data := [][]int{{1, 2, 3, 4, 5}}
	return data
}

func validateUserCode(code string) bool {
	// Define a regular expression pattern for allowed code structure
	pattern := `^\s*(filter|map|sum|min|max|lambda|data|==|!=|<=|>=|<|>|\(|\)|\[|\]|,|\s)*$`
	match, _ := regexp.MatchString(pattern, code)
	return match
}

func inputDataProcessingRules() string {
	// Prompt the user to input data processing rules (Go code).
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter data processing rules: ")
		scanner.Scan()
		userInput := scanner.Text()
		if validateUserCode(userInput) {
			return userInput
		}
		fmt.Println("Invalid code. Please follow the allowed format.")
	}
}

func executeDataProcessing(data [][]int, code string) ([][]int, error) {
	// Check for disallowed functions
	disallowedFunctions := []string{}
	tokens := strings.Fields(code)
	for _, token := range tokens {
		if !allowedFunctions[token] {
			disallowedFunctions = append(disallowedFunctions, token)
		}
	}
	if len(disallowedFunctions) > 0 {
		return nil, fmt.Errorf("Disallowed functions found: %s", strings.Join(disallowedFunctions, ", "))
	}

	cmd := exec.Command("go", "run")
	cmd.Stdin = strings.NewReader(code)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	// Main function to load data, input data processing rules, and execute the code.
	data := loadUserData()
	userCode := inputDataProcessingRules()
	processedData, err := executeDataProcessing(data, userCode)
	if err != nil {
		fmt.Printf("Error executing code: %v\n", err)
		return
	}

	fmt.Println("Processed Data:")
	fmt.Println(processedData)
}
