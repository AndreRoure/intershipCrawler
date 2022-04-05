package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func get(link string, wg sync.WaitGroup) {
	resp, err := http.Get(link)
	if err == nil {
		io.Copy(os.Stdout, resp.Body)
		defer wg.Done()
	}
}
