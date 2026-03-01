package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	prefix := "http://"
	for _, url := range os.Args[1:] {
		urlCopy := url
		if !strings.HasPrefix(urlCopy, prefix) {
			urlCopy = strings.Join([]string{prefix, urlCopy}, "")
		}

		resp, err := http.Get(urlCopy)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
