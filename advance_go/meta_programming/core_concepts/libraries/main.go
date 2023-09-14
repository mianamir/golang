package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

// Configuration represents a configuration object.
type Configuration struct {
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func main() {
	// Load the template
	templateCode, err := ioutil.ReadFile("config_template.go")
	if err != nil {
		panic(err)
	}

	// Parse the template
	fset := token.NewFileSet()
	templateTree, err := parser.ParseFile(fset, "", string(templateCode), parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// Define your configurations
	configurations := []Configuration{
		{DB_NAME: "mydb1", DB_USERNAME: "user1", DB_PASSWORD: "pass1"},
		{DB_NAME: "mydb2", DB_USERNAME: "user2", DB_PASSWORD: "pass2"},
		// Add more configurations as needed
	}

	// Generate configurations
	for _, config := range configurations {
		// Walk through the abstract syntax tree and replace placeholders
		ast.Walk(&replacer{config}, templateTree)
	}

	// Print the variables in the template
	fmt.Println(databaseName) // Will print 'mydb1' for the first configuration
}

// replacer is a visitor that replaces placeholders with actual values.
type replacer struct {
	config Configuration
}

func (r *replacer) Visit(node ast.Node) ast.Visitor {
	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return r
	}

	ident, ok := assignStmt.Lhs[0].(*ast.Ident)
	if !ok {
		return r
	}

	switch ident.Name {
	case "DB_NAME":
		assignStmt.Rhs[0] = &ast.BasicLit{Kind: token.STRING, Value: `"` + r.config.DB_NAME + `"`}
	case "DB_USERNAME":
		assignStmt.Rhs[0] = &ast.BasicLit{Kind: token.STRING, Value: `"` + r.config.DB_USERNAME + `"`}
	case "DB_PASSWORD":
		assignStmt.Rhs[0] = &ast.BasicLit{Kind: token.STRING, Value: `"` + r.config.DB_PASSWORD + `"`}
	}

	return r
}

// Define the variables that correspond to the placeholders
var (
	databaseName   string
	username       string
	password       string
)
