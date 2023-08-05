package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileContent := "File Handling in Golang.."

	file, err := os.Create("./newFile.txt")

	checkErrors(err)

	length, err := io.WriteString(file, fileContent)

	checkErrors(err)

	fmt.Printf("File has been created with %v characters\n", length)

	// now close the file with defer
	defer file.Close()

	// Read Data from Live API

	fmt.Println("Reading JSON and Parsing Data from Live API in Golang")

	const url = "https://mocki.io/v1/d4867d8b-b5d5-4a48-a4ab-79131b5809b8"

	resp, err := http.Get(url)

	checkErrors(err)

	fmt.Printf("Downloaded Response from live API %T: ", resp)

	defer resp.Body.Close()

	// now use this content

	bytes, err := ioutil.ReadAll(resp.Body)

	checkErrors(err)

	content := string(bytes)

	tours := fromJsonFormat(content)

	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

type APILiveData struct {
	Name string
	City string
}

func fromJsonFormat(content string) []APILiveData {
	tours := make([]APILiveData, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()

	checkErrors(err)

	var tour APILiveData

	for decoder.More() {
		err := decoder.Decode(&tour)

		checkErrors(err)

		tours = append(tours, tour)
	}

	return tours

}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	checkErrors(err)

	fmt.Println("Text file: ", string(data))

}
