package main

import (
	"io"
	"net/http"
	"sync"
)

func get(link string, w *sync.WaitGroup, c chan map[string]info) {
	resp, err := http.Get(link)
	if err == nil {
		b, err := io.ReadAll(resp.Body)
		println(link)
		//print(string(b))
		if err == nil {
			if catchCaptcha(string(b)) {
				get(link, w, c)
				return
			}
			w.Add(1)
			go get(next(string(b)), w, c)
			regex(string(b), c)
			defer w.Done()
			return
		}
	}
	defer w.Done()
	//fmt.Println(err)
}
