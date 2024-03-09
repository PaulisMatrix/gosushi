package textsearch

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// document represents a Wikipedia abstract dump document.
type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// loadDocuments loads a Wikipedia abstract dump and returns a slice of documents.
// Dump example: https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
// Compressed version is 120MBs.
// Uncompressed version is 1GBs.
func loadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
