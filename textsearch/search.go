package textsearch

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func TextSearch() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "", "path to the file to read from")
	flag.Parse()

	if filePath == "" {
		fmt.Print("please specify a valid file path...")
		os.Exit(1)
	}

	fmt.Println("please wait while I load the documents....")
	start := time.Now()
	docs, err := loadDocuments(filePath)
	if err != nil {
		log.Fatalf("error: %+v in loading the docs", err)
	}

	fmt.Printf("total time in reading the docs: %f secs\n", time.Now().Sub(start).Seconds())

	fmt.Println("please wait while I index the documents....")
	// generate the index
	start = time.Now()
	invertedIndex := make(Index)
	invertedIndex.Add(docs)
	fmt.Printf("total time in indexing %d docs: %f secs\n", len(docs), time.Now().Sub(start).Seconds())

	fmt.Println("opening up the command prompt, type in your query. type exit to exit the prompt....")

	// search the index
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("exiting due to error: %+v\n", err)
			os.Exit(1)
		}

		text = strings.ToLower(strings.Replace(text, "\n", "", -1))
		if text == "" {
			continue
		}
		if text == "exit" {
			fmt.Print("exiting....")
			return
		}

		searchTime := time.Now()
		matchedIDs := invertedIndex.Search(text)
		fmt.Printf("search found %d docs in %f time (secs)\n", len(matchedIDs), time.Now().Sub(searchTime).Seconds())

		if len(matchedIDs) == 0 {
			fmt.Printf("no docs found matching the query: %s\n", text)
			continue
		}

		for _, id := range matchedIDs {
			doc := docs[id]
			fmt.Printf("doc ID: %d and doc Text: %s\n\n", doc.ID, doc.Text)
		}
	}

}
