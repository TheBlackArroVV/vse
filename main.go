package main

import (
	"fmt"
	"os"
	"path/filepath"
	"vse/index"
	"vse/models"
)

type Query = models.Query

func main() {
	index := index.New("documents")
	indexFolder(index, os.Args[1])

	query := Query{
		Must: []string{os.Args[2]},
	}

	for idx, result := range index.SearchByQuery(query) {
		fmt.Printf("%d: %s", idx, result.Name)
	}
}

func indexFolder(index index.Index, folder string) {
	entires, _ := os.ReadDir(folder)

	for _, entry := range entires {
		if entry.IsDir() {
			indexFolder(index, folder+"/"+entry.Name())
		} else {
			path := filepath.Join(folder, entry.Name())
			file, _ := os.ReadFile(path)
			index.Write(path, string(file))
		}
	}
}
