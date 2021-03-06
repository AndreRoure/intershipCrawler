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
		println(err)
	}
	defer w.Done()
	//fmt.Println(err)
}

//func getDescription(hash string) (description string) {
//	resp, err := http.Get("https://br.indeed.com/?vjk=" + hash)
//	if err == nil {
//		b, err := io.ReadAll(resp.Body)
//		//print(string(b))
//		if err == nil {
//			if catchCaptcha(string(b)) {
//				return getDescription(hash)
//			}
//			return regexDescription(string(b))
//		}
//	}
//	return ""
//}
