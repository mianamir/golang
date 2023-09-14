package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Define your coding standards here
var namingConvention = regexp.MustCompile("^([A-Z][a-zA-Z]*)+$") // Class names must start with an uppercase letter and contain only letters.
var docstringRequired = true                                      // Require methods to have docstrings
var camelCaseVariables = true                                     // Require camel case for variable names
var strReprRequired = true                                         // Require String and Repr methods

// CodeQualityMeta is a struct that holds the coding standards.
type CodeQualityMeta struct{}

// CheckNamingConvention checks if a struct name follows the naming convention.
func (m *CodeQualityMeta) CheckNamingConvention(name string) error {
	if !namingConvention.MatchString(name) {
		return fmt.Errorf("Struct %s does not follow naming convention", name)
	}
	return nil
}

// CheckDocstrings checks if methods have docstrings.
func (m *CodeQualityMeta) CheckDocstrings(name string, s interface{}) error {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		if !strings.HasPrefix(method.Name, "__") {
			doc := method.Func.Type().Out(0).String()
			if docstringRequired && doc != "string" {
				return fmt.Errorf("Method %s in struct %s is missing a docstring", method.Name, name)
			}
		}
	}
	return nil
}

// CheckCamelCaseVariables checks if variable names are in camel case.
func (m *CodeQualityMeta) CheckCamelCaseVariables(name string, s interface{}) error {
	structValue := reflect.ValueOf(s)
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structValue.Type().Field(i).Name
		if camelCaseVariables && !isCamelCase(fieldName) {
			return fmt.Errorf("Variable %s in struct %s should be in camel case", fieldName, name)
		}
	}
	return nil
}

// CheckStrReprMethods checks if the struct has String and Repr methods.
func (m *CodeQualityMeta) CheckStrReprMethods(name string, s interface{}) error {
	structType := reflect.TypeOf(s)
	strMethod, strFound := structType.MethodByName("String")
	reprMethod, reprFound := structType.MethodByName("Repr")

	if strReprRequired && (!strFound || !reprFound) {
		return fmt.Errorf("Struct %s must define both String and Repr methods", name)
	}
	if strReprRequired && strMethod.Func.Type().Out(0).String() != "string" {
		return fmt.Errorf("Method String in struct %s must return a string", name)
	}
	if strReprRequired && reprMethod.Func.Type().Out(0).String() != "string" {
		return fmt.Errorf("Method Repr in struct %s must return a string", name)
	}

	return nil
}

// isCamelCase checks if a string is in camel case.
func isCamelCase(s string) bool {
	if s == "" {
		return true
	}
	runeArray := []rune(s)
	if !isUppercase(runeArray[0]) {
		return false
	}
	prevIsLower := false
	for _, r := range runeArray[1:] {
		if isLowercase(r) {
			prevIsLower = true
		} else if isUppercase(r) {
			if prevIsLower {
				prevIsLower = false
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// isUppercase checks if a rune is an uppercase letter.
func isUppercase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// isLowercase checks if a rune is a lowercase letter.
func isLowercase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func main() {
	type YourStruct struct {
		MyField int
	}

	meta := &CodeQualityMeta{}
	name := "YourStruct"
	err := meta.CheckNamingConvention(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = meta.CheckDocstrings(name, YourStruct{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = meta.CheckCamelCaseVariables(name, YourStruct{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = meta.CheckStrReprMethods(name, YourStruct{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("YourStruct follows coding standards.")
}
