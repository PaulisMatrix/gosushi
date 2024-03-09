package textsearch

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func TextSearch() {
	var filePath, query string
	flag.StringVar(&filePath, "filePath", "", "path to the file to read from")
	flag.StringVar(&query, "query", "", "input query to search")
	flag.Parse()

	start := time.Now()
	docs, err := loadDocuments(filePath)
	if err != nil {
		log.Fatalf("error: %+v in loading the docs", err)
	}

	fmt.Printf("total time in reading the docs: %f\n", time.Now().Sub(start).Seconds())
	fmt.Println("total docs: ", len(docs))

	// generate the index

	// search the index

}
