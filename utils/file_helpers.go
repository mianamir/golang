package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func loadFilesFromFolder(folderPath string, allowedExtensions []string) (map[string][]string, error) {
    loadedData := make(map[string][]string)

    files, err := ioutil.ReadDir(folderPath)
    if err != nil {
        return nil, err
    }

    for _, file := range files {
        if file.IsDir() {
            continue // Skip directories
        }

        ext := filepath.Ext(file.Name())
        ext = strings.TrimPrefix(ext, ".")

        for _, allowedExt := range allowedExtensions {
            if ext == allowedExt {
                filePath := filepath.Join(folderPath, file.Name())

                data, err := readFile(filePath)
                if err != nil {
                    return nil, err
                }

                loadedData[ext] = data
            }
        }
    }

    return loadedData, nil
}

func readFile(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var data []string

    if strings.HasSuffix(filePath, ".csv") {
        csvReader := csv.NewReader(file)

        // Skip the header row if present (you can remove this if you want to keep the header)
        if _, err := csvReader.Read(); err != nil && err != io.EOF {
            return nil, err
        }

        for {
            record, err := csvReader.Read()
            if err == io.EOF {
                break
            }
            if err != nil {
                return nil, err
            }
            data = append(data, strings.Join(record, ","))
        }
    } else {
        // Handle other file types here
        bytes, err := ioutil.ReadAll(file)
        if err != nil {
            return nil, err
        }
        data = strings.Split(string(bytes), "\n")
    }

    return data, nil
}

func main() {
    folderPath := "/path/to/your/files/folder/"
    allowedExtensions := []string{"csv", "txt", "json"} // Specify the allowed file extensions

    loadedData, err := loadFilesFromFolder(folderPath, allowedExtensions)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(loadedData) > 0 {
        for ext, data := range loadedData {
            fmt.Printf("Loaded %s data:\n", strings.ToUpper(ext))
            for _, row := range data {
                fmt.Println(row)
            }
        }
    } else {
        fmt.Println("No data loaded from folder.")
    }
}
