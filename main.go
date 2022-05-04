package main

import (
	"fmt"
	"sync"
)

func request() map[string]info {
	urls := []string{"https://br.indeed.com/empregos?q=R%24%20100.000&l=Bras&vjk=fa673dcce786010c"}
	internships := make(map[string]info)
	messages := make(chan map[string]info)
	routines := sync.WaitGroup{}

	for _, link := range urls {
		routines.Add(1)
		go get(link, &routines, messages)
	}

	go func() {
		fmt.Println("WAITING")
		routines.Wait()
		close(messages)
	}()

	for m := range messages {
		for k, v := range m {
			internships[k] = v
			fmt.Println(k, v)
		}
	}
	fmt.Println("DONE")
	return internships
}

type info struct {
	Titulo string
	Local  string
	Link   string
}
