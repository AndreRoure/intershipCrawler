package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func get(link string, w *sync.WaitGroup) {
	resp, err := http.Get(link)
	if err == nil {
		b, err := io.ReadAll(resp.Body)
		if err == nil {
			regex(string(b))
			defer w.Done()
			return
		}
	}
	defer w.Done()
	fmt.Println(err)
}
