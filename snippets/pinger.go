package main

import (
	"io"
	"log"
	"net/http"
)

func pinger(url string) {

	for {

		go func() {
			var client http.Client
			resp, err := client.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				bytes, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				RespBody := string(bytes)
				log.Println(RespBody)

			}

		}()
	}

}

func main() {

	// ping an url(not recommended tho :p)
	url := "https://www.google.com"
	pinger(url)

}
