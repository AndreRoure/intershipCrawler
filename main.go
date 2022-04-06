package main

import (
	"fmt"
	"sync"
)

func main() {
	urls := []string{"https://br.indeed.com/jobs?q&l=Bras%C3%ADlia%2C%20DF&vjk=bd05e6fb411225f9"}
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
}

type info struct {
	titulo string
	local  string
}
