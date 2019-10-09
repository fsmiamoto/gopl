package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("http://", url) {
			url = fmt.Sprintf("http://%s", url)
		}

		resp, err := http.Get(url)

		if err != nil {
			log.Fatalf("fetch: %v\n", err)
		}

		fmt.Printf("Status: %v", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			log.Fatalf("fetch: reading %s: %v\n", url, err)
		}

	}
}
