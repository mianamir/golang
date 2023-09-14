package main

import (
	"fmt"
)

type Model struct {
    // Define fields with struct tags for validation
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// NewModel initializes a new Model instance
func NewModel(name string, age int) *Model {
    return &Model{Name: name, Age: age}
}

// GetAttr simulates attribute access
func (m *Model) GetAttr(name string) (interface{}, error) {
    switch name {
    case "name":
        return m.Name, nil
    case "age":
        return m.Age, nil
    default:
        return nil, fmt.Errorf("'%s' object has no attribute '%s'", m, name)
    }
}

// SetAttr simulates attribute modification
func (m *Model) SetAttr(name string, value interface{}) error {
    switch name {
    case "name":
        if val, ok := value.(string); ok {
            m.Name = val
            return nil
        }
        return fmt.Errorf("Invalid type for attribute 'name'")
    case "age":
        if val, ok := value.(int); ok {
            m.Age = val
            return nil
        }
        return fmt.Errorf("Invalid type for attribute 'age'")
    default:
        return fmt.Errorf("'%s' object has no attribute '%s'", m, name)
    }
}

func main() {
    // Create a new Model instance
    model := NewModel("John", 30)

    // Access attributes using custom methods
    name, err := model.GetAttr("name")
    if err == nil {
        fmt.Println("Name:", name.(string)) // Cast to the appropriate type
    }

    age, err := model.GetAttr("age")
    if err == nil {
        fmt.Println("Age:", age.(int)) // Cast to the appropriate type
    }

    // Attempt to access a non-existent attribute
    city, err := model.GetAttr("city")
    if err != nil {
        fmt.Println(err) // This will raise an error
    }

    // Modify an attribute using custom methods
    model.SetAttr("name", "Alice")
    fmt.Println("Updated Name:", model.Name)

    // Attempt to modify a non-existent attribute
    err = model.SetAttr("city", "New York")
    if err != nil {
        fmt.Println(err) // This will raise an error
    }
}
