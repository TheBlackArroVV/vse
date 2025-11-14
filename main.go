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
	indexFolder(index, "")

	query := Query{
		Must: []string{"word"},
	}
	fmt.Println(index.SearchByQuery(query))
}

func indexFolder(index index.Index, folder string) {
	entires, _ := os.ReadDir(folder)

	for _, entry := range entires {
		if entry.IsDir() {
			indexFolder(index, folder+"/"+entry.Name())
		} else {
			path := filepath.Join(folder, entry.Name())
			file, _ := os.ReadFile(path)
			index.Write(entry.Name(), string(file))
		}
	}
}
